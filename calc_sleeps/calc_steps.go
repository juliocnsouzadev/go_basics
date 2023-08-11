package main

import (
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) (time.Time, error) {
	date, err := time.Parse(expectedFormat, target)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	daysUntil := time.Until(target).Hours() / 24
	return daysUntil
}
