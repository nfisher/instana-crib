package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/nfisher/instana-crib"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
	"github.com/wcharczuk/go-chart"
)

const (
	// SeriesTimestamp index for the timestamp in the metric results.
	SeriesTimestamp = 0
	// SeriesValue index for the value in the metric results.
	SeriesValue = 1
)

// Exec is the main execution loop of the application.
func Exec(apiToken string, apiURL string, metricName string, pluginType string, queryString string, rollup int64, to int64, windowSize int64) {
	var api instana.InfraQuery
	api, err := instana.NewClient(apiURL, apiToken)
	if err != nil {
		log.Fatalf("unable to create client: %v\n", err)
	}

	metrics, err := api.ListMetrics(queryString, pluginType, []string{metricName}, rollup, windowSize, to)
	if err != nil {
		log.Fatalf("error retrieving metrics: %v\n", err)
	}
	writeCharts(metrics, metricName)

	/*
		snapshots, err := api.ListSnapshots(queryString, pluginType, windowSize)
		if err != nil {
			log.Fatalf("error retrieving snapshots: %v\n", err)
		}
		log.Printf("Snapshots:   %v\n", len(snapshots))
	*/

	log.Printf("Metrics:     %v\n", len(metrics))
}

func main() {
	var apiToken = os.Getenv("INSTANA_TOKEN")
	var apiURL = os.Getenv("INSTANA_URL")

	var metricName string
	var pluginType string
	var queryString string
	var toString string
	var windowString string

	flag.StringVar(&metricName, "metric", "cpu.user", "Metric name to extract")
	flag.StringVar(&queryString, "query", "entity.zone:us-east-2", "Infrastructure query to use as part of the metrics request")
	flag.StringVar(&pluginType, "plugin", "host", "Snapshot plugin type (e.g. host)")
	flag.StringVar(&toString, "to", time.Now().UTC().Format("2006-01-02"), "date time in the format, omitting the clock assumes midnight (YYYY-MM-DD hh:mm:ss)")
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

	log.Printf("API Key Set: %v\n", apiToken != "")
	log.Printf("API URL:     %v\n", apiURL)
	log.Printf("Metric:      %v\n", metricName)
	log.Printf("Query:       %v\n", queryString)
	log.Printf("Rollup:      %v\n", time.Duration(rollup)*time.Second)
	log.Printf("To:          %v\n", toString)
	log.Printf("Window Size: %v\n", time.Duration(windowSize/1000)*time.Second)

	if apiToken == "" {
		log.Fatalln("INSTANA_TOKEN environment variable should be set to the Instana API token. Was a k8s secret created for this?")
	}

	if apiURL == "" {
		log.Fatalln("INSTANA_URL environment variable should be set to the Instana API end-point. Was a k8s secret created for this?")
	}

	to, err := instana.ToInstanaTS(toString)
	if err != nil {
		log.Fatalf("Invalid date time supplied for 'to': %v\n", err)
	}

	Exec(apiToken, apiURL, metricName, pluginType, queryString, rollup, to, windowSize)
}

func writeCharts(metrics []openapi.MetricItem, metricName string) {
	for _, item := range metrics {
		prefix := item.Host
		if prefix == "" {
			prefix = strings.Replace(item.Label, "/", "-", -1)
		}
		//log.Printf("%s %v\n", prefix, item.Metrics[metricName])

		lineChart := newChart(&item, metricName)
		if lineChart == nil {
			continue
		}

		err := renderChart(prefix, lineChart)
		if err != nil {
			log.Printf("error rendering chart %s: %v\n", prefix, err.Error())
		}
	}
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

func renderChart(name string, lineChart *chart.Chart) error {
	buffer := bytes.NewBuffer([]byte{})
	err := lineChart.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%s.png", name))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, buffer)
	if err != nil {
		return err
	}

	return nil
}

func newChart(item *openapi.MetricItem, metricName string) *chart.Chart {
	metricsLen := len(item.Metrics[metricName])
	if metricsLen < 2 {
		log.Printf("no metrics available: %s\n", item.Host)
		return nil
	}
	fmt.Println("len =", metricsLen)
	xValues := make([]float64, metricsLen)
	yValues := make([]float64, metricsLen)

	var metric = item.Metrics[metricName]
	var previous float64
	for i, v := range metric {
		var timestamp = v[SeriesTimestamp]
		var value = v[SeriesValue]
		if math.IsInf(value, 0) || math.IsNaN(value) {
			value = 0.0
		}

		xValues[i] = timestamp
		yValues[i] = value

		if timestamp <= previous {
			log.Printf("warning timestamps out of order %f >= %f: %f %v\n", previous, timestamp, value, int64(v[SeriesTimestamp]))
		}
		previous = timestamp
	}

	title := item.Host
	if title == "" {
		title = item.Label
	}

	return &chart.Chart{
		Title:      title,
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 75,
			},
		},
		XAxis: chart.XAxis{
			Name:           "time",
			NameStyle:      chart.StyleShow(),
			Style:          chart.StyleShow(),
			ValueFormatter: func(v interface{}) string { return time.Unix(int64(v.(float64))/1000, 0).UTC().Format("15:04:05") },
		},
		YAxis: chart.YAxis{
			Name:           metricName,
			NameStyle:      chart.StyleShow(),
			Style:          chart.StyleShow(),
			ValueFormatter: func(v interface{}) string { return chart.FloatValueFormatterWithFormat(v, "%.4f") },
		},
		Width:  800,
		Height: 494,
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}
}
