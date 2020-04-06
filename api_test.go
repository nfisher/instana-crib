package instana_test

import (
	"testing"

	"github.com/nfisher/instana-crib"
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
