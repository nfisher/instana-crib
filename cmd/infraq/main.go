package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/antihax/optional"
	"github.com/instana/instana-crib/pkg/instana/openapi"
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
				WindowSize: 360000,
			},
			Query:   "entity.zone:us-east-2",
			Plugin:  "host",
			Metrics: []string{"cpu.user"},
		}),
	}

	configResp, _, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(ctx, query)
	if err != nil {
		log.Fatalf("error in retrieving metrics: %s\n", err.(openapi.GenericOpenAPIError).Body())
	}

	for _, item := range configResp.Items {
		fmt.Printf("%#v\n", item.Host)
		fmt.Printf("%d\n", len(item.Metrics["cpu.user"]))
	}
}
