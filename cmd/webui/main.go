package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/nfisher/instana-crib"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
)

const (
	// SeriesTimestamp index for the timestamp in the metric results.
	SeriesTimestamp = 0
	// SeriesValue index for the value in the metric results.
	SeriesValue = 1
)

func main() {
	var apiToken = os.Getenv("INSTANA_TOKEN")
	var apiURL = os.Getenv("INSTANA_URL")

	var metricName string
	var pluginType string
	var queryString string
	var toString string
	var windowString string

	flag.StringVar(&metricName, "metric", "cpu.user,cpu.sys", "Metric name to extract")
	flag.StringVar(&queryString, "query", "entity.zone:Instana-Backend", "Infrastructure query to use as part of the metrics request")
	flag.StringVar(&pluginType, "plugin", "host", "Snapshot plugin type (e.g. host)")
	flag.StringVar(&toString, "to", time.Now().UTC().Format("2006-01-02 15:04:05"), "date time in the format, omitting the clock assumes midnight (YYYY-MM-DD hh:mm:ss)")
	flag.StringVar(&windowString, "window", "60s", `metric window size (valid time units are "s", "m", "h")`)

	flag.Parse()

	windowSize, err := instana.ParseDuration(windowString)
	if err != nil {
		log.Fatalln(err)
	}

	rollup, err := rollupForWindow(windowSize)
	if err != nil {
		log.Fatalln(err)
	}

	if apiToken == "" {
		log.Fatalln("INSTANA_TOKEN environment variable should be set to the Instana API token. Was a k8s secret created for this?")
	}

	if apiURL == "" {
		log.Fatalln("INSTANA_URL environment variable should be set to the Instana API end-point. Was a k8s secret created for this?")
	}

	api, err := instana.NewClient(apiURL, apiToken)
	if err != nil {
		log.Fatalf("unable to create client: %v\n", err)
	}

	var metricValue atomic.Value
	metricValue.Store([]openapi.MetricItem{})

	go func(api instana.InfraQuery) {
		for {
			toString := time.Now().UTC().Format("2006-01-02 15:04:05")
			to, err := instana.ToInstanaTS(toString)
			if err != nil {
				log.Printf("Invalid date time supplied for 'to': %v\n", err)
			}

			metrics, err := api.ListMetrics(queryString, pluginType, strings.Split(metricName, ","), rollup, windowSize, to)
			if err != nil {
				log.Printf(err.Error())
			} else {
				// TODO: replace else condition "happy path" with more idiomatic go.
				metricValue.Store(metrics)
			}

			// TODO: adjust sleep for true 1s intervals
			time.Sleep(1 * time.Second)
		}
	}(api)

	var reMetricName = regexp.MustCompile(`^[a-zA-Z0-9.]+$`)

	http.HandleFunc("/heatmap_data", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "text/csv")
		req.ParseForm()
		metricName := req.Form.Get("metric")
		if !reMetricName.MatchString(metricName) {
			http.Error(w, "Invalid metric name", http.StatusBadRequest)
			return
		}

		metrics := metricValue.Load().([]openapi.MetricItem)
		hist := instana.ToPercentageHeatmap(metrics, metricName)
		tab := instana.ToTabular(hist)
		enc := csv.NewWriter(w)
		err = enc.WriteAll(tab)
		if err != nil {
			http.Error(w, fmt.Sprintf("error encoding metrics: %v", err), http.StatusInternalServerError)
			return
		}
		enc.Flush()
	})

	http.HandleFunc("/heatmap", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "text/html")
		w.Write([]byte(heatmapHtml))
	})

	http.ListenAndServe(":8000", nil)
}

type MetricResponse struct {
	Hosts []string
	Headers []string
	Data [][]*float32
}

func rollupForWindow(windowSize int64) (int64, error) {
	rollup := windowSize / 1000 / 600
	if rollup <= 1 {
		return 1, nil
	} else if rollup <= 5 {
		return 5, nil
	} else if rollup <= 60 {
		return 60, nil
	} else if rollup <= 300 {
		return 300, nil
	} else if rollup <= 3600 {
		return 3600, nil
	}

	return 0, errors.New("rollup is too large for API call, maximum call size is 25 days")
}

const heatmapHtml = `<!DOCTYPE html>
<meta charset="utf-8">
<link href="https://fonts.googleapis.com/css2?family=Work+Sans&display=swap" rel="stylesheet">
<title>Host Monitor</title>
<style>
html {
  font-family: 'Work Sans', sans-serif;
}
h2 {
    margin-bottom: 0.1em;
	margin-top: 0.2em;
}
</style>
<script src="https://d3js.org/d3.v4.js"></script>

<h2>CPU User (<span id="g_cpu_user_count"></span> hosts)</h2>
<div id="g_cpu_user"></div>

<h2>CPU System (<span id="g_cpu_sys_count"></span> hosts)</h2>
<div id="g_cpu_sys"></div>

<script>

function drawGraph() {
    var w=window,
    d=document,
    e=d.documentElement,
    g=d.getElementsByTagName('body')[0],
    x=w.innerWidth||e.clientWidth||g.clientWidth;

    // set the dimensions and margins of the graph
    var margin = {top: 5, right: 30, bottom: 55, left: 35},
      width = x - margin.left - margin.right,
      height = 150 - margin.top - margin.bottom;
    
    //Read the data
    d3.csv("/heatmap_data?metric=cpu.user", function(data) {
    var svg = d3.select("#g_cpu_user")
        .select("svg")
        .remove();
    // append the svg object to the body of the page
    var svg = d3.select("#g_cpu_user")
    .append("svg")
      .attr("width", width + margin.left + margin.right)
      .attr("height", height + margin.top + margin.bottom)
    .append("g")
      .attr("transform",
            "translate(" + margin.left + "," + margin.top + ")");
      var groups = {};
      var subGroups = {};
      var upper = 0;
	  for (var i = 0; i < data.length; i++) {
		var g = data[i]["group"];
		var s = data[i]["variable"];
		groups[g] = true;
		subGroups[s] = true;
      }
      for (var i = 0; i < Object.keys(subGroups).length; i++) {
		upper += parseInt(data[i]["value"]);
      }
      // Labels of row and columns
      var myGroups = Object.keys(groups);
      var myVars = Object.keys(subGroups);
      
      // Build X scales and axis:
      var x = d3.scaleBand()
        .range([ 0, width])
        .domain(myGroups)
        .padding(0.01);
      svg.append("g")
        .attr("transform", "translate(0," + height + ")")
        .call(d3.axisBottom(x).tickValues(x.domain().filter(function(d,i){ return !(i%2)})))
		.selectAll("text") 
			.style("text-anchor", "end")
			.attr("dx", "-.8em")
			.attr("dy", ".15em")
			.attr("transform", "rotate(-65)");
      
      // Build X scales and axis:
      var y = d3.scaleBand()
        .range([ height, 0 ])
        .domain(myVars)
        .padding(0.01);
      svg.append("g")
        .call(d3.axisLeft(y).tickValues(y.domain().filter(function(d,i){ return !(i%2)})));

		d3.select("#g_cpu_user_count")
			.text("" + upper);
      // Build color scale
      var myColor = d3.scaleLinear()
        .range(["white", "#ddd", "#333"])
        .domain([0, 1, upper])
        svg.selectAll()
          .data(data, function(d) {return d.group+':'+d.variable;})
          .enter()
          .append("rect")
          .attr("x", function(d) { return x(d.group) })
          .attr("y", function(d) { return y(d.variable) })
          .attr("width", x.bandwidth() )
          .attr("height", y.bandwidth() )
          .style("fill", function(d) { return myColor(d.value)} )
    })
}
function drawGraph2() {
    var w=window,
    d=document,
    e=d.documentElement,
    g=d.getElementsByTagName('body')[0],
    x=w.innerWidth||e.clientWidth||g.clientWidth;

    // set the dimensions and margins of the graph
    var margin = {top: 5, right: 30, bottom: 55, left: 35},
      width = x - margin.left - margin.right,
      height = 150 - margin.top - margin.bottom;
    
    //Read the data
    d3.csv("/heatmap_data?metric=cpu.sys", function(data) {
    var svg = d3.select("#g_cpu_sys")
        .select("svg")
        .remove();
    // append the svg object to the body of the page
    var svg = d3.select("#g_cpu_sys")
    .append("svg")
      .attr("width", width + margin.left + margin.right)
      .attr("height", height + margin.top + margin.bottom)
    .append("g")
      .attr("transform",
            "translate(" + margin.left + "," + margin.top + ")");
      var groups = {};
      var subGroups = {};
      var upper = 0;
	  for (var i = 0; i < data.length; i++) {
		var g = data[i]["group"];
		var s = data[i]["variable"];
		groups[g] = true;
		subGroups[s] = true;
      }
      for (var i = 0; i < Object.keys(subGroups).length; i++) {
		upper += parseInt(data[i]["value"]);
      }
      // Labels of row and columns
      var myGroups = Object.keys(groups);
      var myVars = Object.keys(subGroups);
      
      // Build X scales and axis:
      var x = d3.scaleBand()
        .range([ 0, width])
        .domain(myGroups)
        .padding(0.01);
      svg.append("g")
        .attr("transform", "translate(0," + height + ")")
        .call(d3.axisBottom(x).tickValues(x.domain().filter(function(d,i){ return !(i%2)})))
		.selectAll("text") 
			.style("text-anchor", "end")
			.attr("dx", "-.8em")
			.attr("dy", ".15em")
			.attr("transform", "rotate(-65)");
      
      // Build X scales and axis:
      var y = d3.scaleBand()
        .range([ height, 0 ])
        .domain(myVars)
        .padding(0.01);
      svg.append("g")
        .call(d3.axisLeft(y).tickValues(y.domain().filter(function(d,i){ return !(i%2)})));
      
		d3.select("#g_cpu_sys_count")
			.text("" + upper);
      // Build color scale
      // Build color scale
      var myColor = d3.scaleLinear()
        .range(["white", "#ddd", "#333"])
        .domain([0, 1, upper])
        svg.selectAll()
          .data(data, function(d) {return d.group+':'+d.variable;})
          .enter()
          .append("rect")
          .attr("x", function(d) { return x(d.group) })
          .attr("y", function(d) { return y(d.variable) })
          .attr("width", x.bandwidth() )
          .attr("height", y.bandwidth() )
          .style("fill", function(d) { return myColor(d.value)} )
    })
}
window.addEventListener('resize', drawGraph);
window.addEventListener('resize', drawGraph2);
drawGraph();
drawGraph2();
setInterval(drawGraph, 500);
setInterval(drawGraph2, 500);
</script>`