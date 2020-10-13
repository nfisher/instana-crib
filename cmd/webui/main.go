package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
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

type Timeseries struct {
	Values []float64 `json:"values"`
}

func main() {
	var apiToken = os.Getenv("INSTANA_TOKEN")
	var apiURL = os.Getenv("INSTANA_URL")

	var windowString string

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

	log.Println("URL:", apiURL)

	api, err := instana.NewClient(apiURL, apiToken)
	if err != nil {
		log.Fatalf("unable to create client: %v\n", err)
	}

	var metricValue atomic.Value
	m := map[string][]openapi.MetricItem{
		"host": {},
	}
	metricValue.Store(m)

	go func(api instana.InfraQuery) {
		c := time.Tick(1 * time.Second)
		for next := range c {
			_ = next
			toString := time.Now().UTC().Format("2006-01-02 15:04:05")
			to, err := instana.ToInstanaTS(toString)
			if err != nil {
				log.Printf("Invalid date time supplied for 'to': %v\n", err)
			}

			////
			//
			// == acceptor ==
			// "plugin":"dropwizardApplicationContainer",
			// "query":"appdata-writer",
			//
			// metrics.guage.KPI.outgoing.spans.error_rate
			//
			// == ad-writer ==
			// metrics.gauges.KPI.incoming.raw_spans.error_rate
			//
			adWriterMetrics, err := api.ListMetrics(
				"entity.label:*appdata-writer*",
				"dropwizardApplicationContainer",
				[]string{
					"metrics.gauges.KPI.incoming.raw_spans.error_rate",
					"metrics.meters.KPI.incoming.raw_spans.calls",
				},
				rollup,
				windowSize,
				to)
			if err != nil {
				log.Printf(err.Error())
			}

			// == ad-processor ==
			// "plugin":"dropwizardApplicationContainer",
			// "query":"appdata-processor",
			//
			// metrics.gauges.KPI.incoming.span_messages.error_rate
			//
			adProcessorMetrics, err := api.ListMetrics(
				"entity.label:*appdata-processor*",
				"dropwizardApplicationContainer",
				[]string{
					"metrics.gauges.KPI.incoming.span_messages.error_rate",
					"metrics.meters.KPI.incoming.span_messages.calls",
				},
				rollup,
				windowSize,
				to)
			if err != nil {
				log.Printf(err.Error())
			}

			// == filler ==
			// "plugin":"dropwizardApplicationContainer",
			// "query":"filler",
			// metrics.guage.KPI.incoming.raw_messages.error_rate
			//
			fillerMetrics, err := api.ListMetrics(
				"entity.label:filler*",
				"dropwizardApplicationContainer",
				[]string{
					"metrics.gauges.KPI.incoming.raw_messages.error_rate",
					"metrics.gauges.com.instana.filler.service.snapshot.OnlineSnapshotsLimit.online-snapshots-count",
				},
				rollup,
				windowSize,
				to)
			if err != nil {
				log.Printf(err.Error())
			}

			hostMetrics, err := api.ListMetrics(
				"entity.type:host AND entity.zone:Instana-*",
				"host",
				[]string{ "cpu.user", "cpu.sys", "cpu.wait" },
				rollup,
				windowSize,
				to)
			if err != nil {
				log.Printf(err.Error())
			}

			m := map[string][]openapi.MetricItem{
				"appdataProcessor": adProcessorMetrics,
				"appdataWriter": adWriterMetrics,
				"filler": fillerMetrics,
				"host":   hostMetrics,
			}

			metricValue.Store(m)
		}
	}(api)

	var reMetricName = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

	http.HandleFunc("/ts_sum", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing submitted values: %v", err), http.StatusInternalServerError)
			return
		}

		entityName := req.Form.Get("entity")
		metricName := req.Form.Get("metric")

		if !reMetricName.MatchString(metricName) {
			http.Error(w, "invalid metric name", http.StatusBadRequest)
			return
		}

		metrics := metricValue.Load().(map[string][]openapi.MetricItem)
		metric, ok := metrics[entityName]
		if !ok {
			http.Error(w, "invalid entity name", http.StatusBadRequest)
			return
		}

		values := instana.Sum(metric, metricName)
		ts := Timeseries {
			Values: values,
		}

		w.Header().Set("Content-type", "text/csv")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		err = json.NewEncoder(gz).Encode(&ts)
		if err != nil {
			http.Error(w, fmt.Sprintf("error json encoding values: %v", err), http.StatusInternalServerError)
			return
		}
	})


	http.HandleFunc("/heatmap_data", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing submitted values: %v", err), http.StatusInternalServerError)
			return
		}

		entityName := req.Form.Get("entity")
		metricName := req.Form.Get("metric")

		if !reMetricName.MatchString(metricName) {
			http.Error(w, "invalid metric name", http.StatusBadRequest)
			return
		}

		metrics := metricValue.Load().(map[string][]openapi.MetricItem)
		metric, ok := metrics[entityName]
		if !ok {
			http.Error(w, "invalid entity name", http.StatusBadRequest)
			return
		}

		hist := instana.ToPercentageHeatmap(metric, metricName)
		tab := instana.ToTabular(hist)
		w.Header().Set("Content-type", "text/csv")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		enc := csv.NewWriter(gz)
		err = enc.WriteAll(tab)
		if err != nil {
			http.Error(w, fmt.Sprintf("error encoding metrics: %v", err), http.StatusInternalServerError)
			return
		}
		enc.Flush()
	})

	http.Handle("/", http.FileServer(http.Dir("./html")))
	log.Println("binding to :8000")
	http.ListenAndServe(":8000", nil)
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
