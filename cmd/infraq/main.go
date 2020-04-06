package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
	"github.com/wcharczuk/go-chart"
)

const (
	// SeriesTimestamp index for the timestamp in the metric results.
	SeriesTimestamp = 0
	// SeriesValue index for the value in the metric results.
	SeriesValue = 1
)

func newConfiguration(apiURL string, isInsecure bool) (*openapi.Configuration, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			// ignore expired SSL certificates
			TLSClientConfig: &tls.Config{InsecureSkipVerify: isInsecure},
		},
	}

	configuration := openapi.NewConfiguration()
	configuration.BasePath = apiURL
	configuration.Host = u.Hostname()
	configuration.HTTPClient = httpClient

	return configuration, nil
}

func main() {
	var apiToken = os.Getenv("INSTANA_TOKEN")
	var apiURL = os.Getenv("INSTANA_URL")

	var metricName string
	var queryString string
	var pluginType string
	var windowSize int64

	flag.StringVar(&metricName, "metric", "cpu.user", "Metric name to extract")
	flag.StringVar(&queryString, "query", "entity.zone:us-east-2", "Infrastructure query to use as part of the metrics request")
	flag.StringVar(&pluginType, "plugin", "host", "Snapshot plugin type (e.g. host)")
	flag.Int64Var(&windowSize, "window", 3600, "metric window size in seconds")
	flag.Parse()

	rollup := windowSize / 1000 / 600
	if rollup <= 1 {
		rollup = 1
	} else if rollup <= 5 {
		rollup = 5
	} else if rollup <= 60 {
		rollup = 60
	} else if rollup <= 300 {
		rollup = 300
	} else if rollup <= 3600 {
		rollup = 3600
	} else {
		log.Fatal("Rollup is too large for this API call, maximum is 25 days (2,160,000,000)")
	}

	log.Printf("API Key Set: %v\n", apiToken != "")
	log.Printf("API URL:     %v\n", apiURL)
	log.Printf("Metric:      %v\n", metricName)
	log.Printf("Query:       %v\n", queryString)
	log.Printf("Rollup:      %v\n", rollup)
	log.Printf("Window Size: %v\n", time.Duration(windowSize/1000)*time.Second)

	if apiToken == "" {
		panic("INSTANA_TOKEN environment variable should be set to the Instana API token. Was a k8s secret created for this?")
	}

	if apiURL == "" {
		panic("INSTANA_URL environment variable should be set to the Instana API end-point. Was a k8s secret created for this?")
	}

	configuration, err := newConfiguration(apiURL, true)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := openapi.NewAPIClient(configuration)
	ctx := context.WithValue(
		context.Background(),
		openapi.ContextAPIKey,
		openapi.APIKey{
			Key:    apiToken,
			Prefix: "apiToken",
		})

	var query = &openapi.GetInfrastructureMetricsOpts{
		GetCombinedMetrics: optional.NewInterface(openapi.GetCombinedMetrics{
			TimeFrame: openapi.TimeFrame{
				WindowSize: windowSize,
			},
			Rollup:  int32(rollup),
			Query:   queryString,
			Plugin:  pluginType,
			Metrics: []string{metricName},
		}),
	}

	configResp, httpResp, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(ctx, query)
	if err != nil {
		log.Fatalf("error in retrieving metrics: %s\n", err.(openapi.GenericOpenAPIError).Body())
	}

	log.Printf("Limit Remaining: %v\n", httpResp.Header.Get("X-Ratelimit-Remaining"))

	if len(configResp.Items) < 1 {
		log.Fatalln("No metrics found")
	}

	for _, item := range configResp.Items {
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
			Name:      "time",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
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
