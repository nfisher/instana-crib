package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/antihax/optional"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
	"github.com/wcharczuk/go-chart"
)

const (
	// SeriesTimestamp index for the timestamp in the metric results.
	SeriesTimestamp = 0
	// SeriesValue index for the value in the metric results.
	SeriesValue = 1
	// Metric is the metric.
	Metric = "cpu.user"
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
	var queryString string

	flag.StringVar(&queryString, "query", "entity.zone:us-east-2", "Infrastructure query to use as part of the metrics request")
	flag.Parse()

	log.Printf("API Key Set: %v\n", apiToken != "")
	log.Printf("API URL:     %v\n", apiURL)

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
				WindowSize: 3600,
			},
			Rollup:  1,
			Query:   queryString,
			Plugin:  "host",
			Metrics: []string{Metric},
		}),
	}

	configResp, httpResp, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(ctx, query)
	if err != nil {
		log.Fatalf("error in retrieving metrics: %s\n", err.(openapi.GenericOpenAPIError).Body())
	}

	log.Printf("Rate Limit Remaining: %#v\n", httpResp.Header.Get("X-Ratelimit-Remaining"))

	if len(configResp.Items) < 1 {
		log.Fatalln("No metrics found")
	}

	for _, item := range configResp.Items {
		prefix := item.Host
		log.Println(prefix, len(item.Metrics[Metric]))

		lineChart := newChart(&item)
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

func newChart(item *openapi.MetricItem) *chart.Chart {
	metricsLen := len(item.Metrics[Metric])
	if metricsLen < 2 {
		log.Printf("no metrics available: %s\n", item.Host)
		return nil
	}
	xValues := make([]float64, metricsLen, metricsLen)
	yValues := make([]float64, metricsLen, metricsLen)

	var lastTimestamp float32
	for i, v := range item.Metrics[Metric] {
		currentTimestamp := v[SeriesTimestamp]
		value := v[SeriesValue]

		if lastTimestamp >= currentTimestamp {
			log.Printf("timestamps out of order %f, %f\n", lastTimestamp, currentTimestamp)
		}
		lastTimestamp = currentTimestamp

		xValues[i] = float64(currentTimestamp)
		yValues[i] = float64(value)
	}

	return &chart.Chart{
		Title:      item.Host,
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
			Name:      "usage",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Width:  800,
		Height: 600,
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}
}
