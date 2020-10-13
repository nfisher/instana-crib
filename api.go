package instana

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
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

	remaining, err := strconv.ParseInt(httpResp.Header.Get("X-Ratelimit-Remaining"), 10, 32)
	if err != nil {
		log.Println("unable to convert remaining rate limit")
	}
	if remaining < 25 && remaining % 5 == 0 {
		log.Printf("warning minimal requests remaining: %v\n", remaining)
	}

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
		gerr, ok := err.(openapi.GenericOpenAPIError)
		if !ok {
			return nil, fmt.Errorf("error retrieving metrics: %v", err)
		}
		return nil, fmt.Errorf("error in retrieving metrics: %s", gerr.Body())
	}

	remaining, err := strconv.ParseInt(httpResp.Header.Get("X-Ratelimit-Remaining"), 10, 32)
	if err != nil {
		log.Println("unable to convert remaining rate limit")
	}
	if remaining < 25 && remaining % 5 == 0 {
		log.Printf("Warning Requests Remaining: %v\n", httpResp.Header.Get("X-Ratelimit-Remaining"))
	}
	if len(metricsResp.Items) < 1 {
		return nil, errors.New("no metrics found")
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

const percentBuckets = 21

type PercentageHeatmap map[string][percentBuckets]int

const hoursMinutesSeconds = "15:04:05"

func Sum(items []openapi.MetricItem, metric string) []float64 {
	series := make(map[string]float64)
	for _, item := range items {
		ts := item.Metrics[metric]
		for _, m := range ts {
			t := time.Unix(int64(m[0]/1000), 0).UTC().Format(hoursMinutesSeconds)
			v := series[t]
			v += m[1]
			series[t] = v
		}
	}
	var keys []string
	for k := range series {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var ts []float64
	for _, k := range keys {
		ts = append(ts, series[k])
	}
	return ts
}

func ToPercentageHeatmap(items []openapi.MetricItem, metric string) PercentageHeatmap {
	var ph = make(PercentageHeatmap)
	for _, item := range items {
		ts := item.Metrics[metric]
		for _, m := range ts {
			t := time.Unix(int64(m[0]/1000), 0).UTC().Format(hoursMinutesSeconds)
			v := int(math.Floor(m[1] * percentBuckets)) // scale to an index
			if v == 0 && m[1] > 0 {
				v = 1
			}
			if v > percentBuckets - 1 {
				v = percentBuckets - 1
			}
			hist, ok := ph[t]
			if !ok {
				hist = [percentBuckets]int{}
			}
			hist[v]++
			ph[t] = hist
		}
	}
	return ph
}

func ToTabular(hist PercentageHeatmap) [][]string {
	var tab = [][]string{{"group","variable","value"}}
	var labels []string
	for l := range hist {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	for _, l := range labels {
		for i, v := range hist[l] {
			var p string
			if i == 0 {
				p = "0%"
			} else {
				p = fmt.Sprintf("%d%%", i * 100 / (percentBuckets -1))
			}
			s := fmt.Sprintf("%d", v)
			tab = append(tab, []string{l, p, s})
		}
	}
	return tab
}
