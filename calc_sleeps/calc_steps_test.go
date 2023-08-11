package main

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		input       string
		expected    time.Time
		shouldError bool
	}{
		{"2022-01-01", time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"2022-02-30", time.Time{}, true},
		{"2022-13-01", time.Time{}, true},
		{"2022-01-01T00:00:00Z", time.Time{}, true},
	}

	for _, tc := range testCases {
		actual, err := parseTime(tc.input)
		if tc.shouldError {
			if err == nil {
				t.Errorf("parseTime(%q) should have returned an error", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("parseTime(%q) returned an error: %v", tc.input, err)
			}
			if !actual.Equal(tc.expected) {
				t.Errorf("parseTime(%q) = %v, expected %v", tc.input, actual, tc.expected)
			}
		}
	}
}

func TestCalcSleeps(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		target   time.Time
		expected float64
	}{
		{now.Add(time.Hour * 24), 1},
		{now.Add(time.Hour * 48), 2},
		{now.Add(time.Hour * 24 * 365), 365},
	}

	for _, tc := range testCases {
		actual := calcSleeps(tc.target)
		if actual != tc.expected {
			t.Errorf("calcSleeps(%v) = %v, expected %v", tc.target, actual, tc.expected)
		}
	}
}
