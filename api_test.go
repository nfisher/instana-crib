package instana_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nfisher/instana-crib"
	"github.com/nfisher/instana-crib/pkg/instana/openapi"
)

func Test_ToInstanaTS(t *testing.T) {
	t.Parallel()

	td := []struct {
		name     string
		input    string
		expected int64
		hasError bool
	}{
		{"valid date time", "2020-04-06 00:00:01", 1586131201 * 1000, false},
		{"default time to midnight", "2020-04-06", 1586131200 * 1000, false},
		{"invalid", "adbdcadsca1234", -1, true},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := instana.ToInstanaTS(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("ToInstana(%v) error = %v, got %v, want %v", tc.input, err, err != nil, tc.hasError)
			}

			if actual != tc.expected {
				t.Errorf("ToInstana(%v) = %v, want %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func Test_InstanaDuration(t *testing.T) {
	t.Parallel()

	td := []struct {
		name     string
		input    string
		expected int64
		hasError bool
	}{
		{"secound", "1s", 1000, false},
		{"minute", "1m", 60 * 1000, false},
		{"hour", "1h", 60 * 60 * 1000, false},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := instana.ParseDuration(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("ToInstana(%v) error = %v, got %v, want %v", tc.input, err, err != nil, tc.hasError)
			}

			if actual != tc.expected {
				t.Errorf("ToInstana(%v) = %v, want %v", tc.input, actual, tc.expected)
			}
		})
	}
}

const CpuUser = "cpu.user"

func cpuUser(startEpoch int64, values []float64) map[string][][]float64 {
	var ts [][]float64
	for i, v := range values {
		t := float64(int64(i) + startEpoch) * 1000
		ts = append(ts, []float64{t, v})
	}
	m := map[string][][]float64{
		CpuUser: ts,
	}
	return m
}

func Test_Sum_aligns_seconds(t *testing.T) {
	input := []openapi.MetricItem{
		{Metrics: cpuUser(1601553600, []float64{0.01, 0.01, 0.01})},
		{Metrics: cpuUser(1601553601, []float64{0.01, 0.02})},
	}
	expected := []float64{0.01, 0.02, 0.03}

	actual := instana.Sum(input, CpuUser)

	if !cmp.Equal(actual, expected) {
		t.Errorf("Sum() -got/+want:\n%s", cmp.Diff(expected, actual))
	}
}

const percentBucketSize = 21

func Test_ToPercentageHeatmap(t *testing.T) {
	td := map[string]struct {
		input []openapi.MetricItem
		expected instana.PercentageHeatmap
	}{
		"multiple moments": {
			[]openapi.MetricItem{{Metrics: cpuUser(1601553600, []float64{0, 0.01, 0.1})}},
			instana.PercentageHeatmap{"12:00:00": [percentBucketSize]int{1}, "12:00:01": [percentBucketSize]int{0, 1}, "12:00:02": [percentBucketSize]int{0, 0, 1}}},
		"multiple items": {
			[]openapi.MetricItem{
				{Metrics: cpuUser(1601553600, []float64{0.01, 0.01, 0.01})},
				{Metrics: cpuUser(1601553600, []float64{0.1, 0.01, 1.0})}},
			instana.PercentageHeatmap{
				"12:00:00": [percentBucketSize]int{0, 1, 1},
				"12:00:01": [percentBucketSize]int{0, 2},
				"12:00:02": [percentBucketSize]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}},
	}

	for name, tc := range td {
		t.Run(name, func(t *testing.T) {
			ph := instana.ToPercentageHeatmap(tc.input, CpuUser)
			if !cmp.Equal(ph, tc.expected) {
				t.Errorf("ToPercentageHeatmap(%#v) -got/+want:\n%s", tc.input, cmp.Diff(tc.expected, ph))
			}
		})
	}
}

func Test_ToTabular(t *testing.T) {
	hist := instana.PercentageHeatmap{"12:00:00": [percentBucketSize]int{1, 1}, "12:00:01": [percentBucketSize]int{2}, "12:00:02": [percentBucketSize]int{1, 0,  0,  0,  0,  0,  0,  0,  0, 1}}
	tab := instana.ToTabular(hist)
	expected := [][]string{
		{"group","variable","value"},
		{"12:00:00", "0%", "1"},
		{"12:00:00", "5%", "1"},
		{"12:00:00", "10%", "0"},
		{"12:00:00", "15%", "0"},
		{"12:00:00", "20%", "0"},
		{"12:00:00", "25%", "0"},
		{"12:00:00", "30%", "0"},
		{"12:00:00", "35%", "0"},
		{"12:00:00", "40%", "0"},
		{"12:00:00", "45%", "0"},
		{"12:00:00", "50%", "0"},
		{"12:00:00", "55%", "0"},
		{"12:00:00", "60%", "0"},
		{"12:00:00", "65%", "0"},
		{"12:00:00", "70%", "0"},
		{"12:00:00", "75%", "0"},
		{"12:00:00", "80%", "0"},
		{"12:00:00", "85%", "0"},
		{"12:00:00", "90%", "0"},
		{"12:00:00", "95%", "0"},
		{"12:00:00", "100%", "0"},
		{"12:00:01", "0%", "2"},
	}
	if !cmp.Equal(tab[:23], expected) {
		t.Errorf("-got/+want:\n%s", cmp.Diff(expected, tab[:22]))
	}
}
