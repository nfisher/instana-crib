package instana

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/antihax/optional"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
)

// InfraQuery is a common interface for infrastructure queries.
type InfraQuery interface {
	ListMetrics(queryString string, pluginType string, metrics []string, rollup int64, windowSize int64, to int64) ([]openapi.MetricItem, error)
	ListSnapshots(queryString string, pluginType string, windowSize int64) ([]openapi.SnapshotItem, error)
}

// NewClient builds an Instana API client from the specified URL and token.
func NewClient(apiURL string, apiToken string) (InfraQuery, error) {
	configuration, err := newConfiguration(apiURL, true)
	if err != nil {
		return nil, err
	}

	client := openapi.NewAPIClient(configuration)
	ctx := context.WithValue(
		context.Background(),
		openapi.ContextAPIKey,
		openapi.APIKey{
			Key:    apiToken,
			Prefix: "apiToken",
		},
	)

	api := &InfraQueryAPI{
		client: client,
		ctx:    ctx,
	}

	return api, nil
}

// InfraQueryAPI is a concrete implementation fo the InfraQuery interface using the openapi client.
type InfraQueryAPI struct {
	client *openapi.APIClient
	ctx    context.Context
}

// ListSnapshots returns the list of snapshots matching the supplied query parameters.
func (api *InfraQueryAPI) ListSnapshots(queryString string, pluginType string, windowSize int64) ([]openapi.SnapshotItem, error) {
	var snapshotsQuery = &openapi.GetSnapshotsOpts{
		Query:      optional.NewString(queryString),
		Plugin:     optional.NewString(pluginType),
		WindowSize: optional.NewInt64(windowSize),
	}
	snapshotResp, httpResp, err := api.client.InfrastructureMetricsApi.GetSnapshots(api.ctx, snapshotsQuery)
	if err != nil {
		log.Fatalf("error in retrieving snapshots: %s\n", err.(openapi.GenericOpenAPIError).Body())
	}

	log.Printf("Remaining:   %v\n", httpResp.Header.Get("X-Ratelimit-Remaining"))

	return snapshotResp.Items, nil
}

// ListMetrics returns the list of metrics matching the supplied query parameters.
func (api *InfraQueryAPI) ListMetrics(queryString string, pluginType string, metrics []string, rollup int64, windowSize int64, to int64) ([]openapi.MetricItem, error) {
	var query = &openapi.GetInfrastructureMetricsOpts{
		GetCombinedMetrics: optional.NewInterface(openapi.GetCombinedMetrics{
			TimeFrame: openapi.TimeFrame{
				WindowSize: windowSize,
				To:         to,
			},
			Rollup:  int32(rollup),
			Query:   queryString,
			Plugin:  pluginType,
			Metrics: metrics,
		}),
	}

	metricsResp, httpResp, err := api.client.InfrastructureMetricsApi.GetInfrastructureMetrics(api.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving metrics: %s", err.(openapi.GenericOpenAPIError).Body())
	}

	log.Printf("Remaining:   %v\n", httpResp.Header.Get("X-Ratelimit-Remaining"))

	if len(metricsResp.Items) < 1 {
		return nil, errors.New("No metrics found")
	}

	return metricsResp.Items, nil
}

// ToInstanaTS converts a datetime string to an instana Dynamic Focus Query timestamp.
func ToInstanaTS(datetime string) (int64, error) {
	if len(datetime) == 10 {
		datetime += " 00:00:00"
	}

	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return -1, err
	}
	return t.Unix() * 1000, nil
}

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

// ParseDuration parses a duration string and scales it to the value expected by the Instana API.
func ParseDuration(s string) (int64, error) {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return -1, err
	}
	return int64(duration / time.Millisecond), nil
}
