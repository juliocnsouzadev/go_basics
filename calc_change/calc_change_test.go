package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateChange(t *testing.T) {
	testCases := []struct {
		amount   float64
		expected map[coin]int
	}{
		{0.0, map[coin]int{}},
		{0.01, map[coin]int{OnePenny: 1}},
		{0.05, map[coin]int{FivePence: 1}},
		{0.10, map[coin]int{TenPence: 1}},
		{0.25, map[coin]int{TwentyPence: 1, FivePence: 1}},
		{0.99, map[coin]int{FiftyPence: 1, TwentyPence: 2, FivePence: 1, OnePenny: 4}},
		{1.00, map[coin]int{OnePound: 1}},
		{2.00, map[coin]int{OnePound: 2}},
		{5.00, map[coin]int{OnePound: 5}},
		{10.00, map[coin]int{OnePound: 10}},
	}

	for _, tc := range testCases {
		actual := calculateChange(tc.amount)
		require.NoError(t, Compare(actual, tc.expected))
	}

}

func Compare(actual, expected map[coin]int) error {
	if len(actual) != len(expected) {
		return fmt.Errorf("lengths mismatch -> expected %v, actual %v", len(expected), len(actual))
	}
	for k, v := range actual {
		if expected[k] != v {
			return fmt.Errorf("nb of coins missmatch for value %v -> expected %v, actual %v", k, expected[k], v)
		}
	}
	return nil
}
