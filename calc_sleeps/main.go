package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	date, err := time.Parse(expectedFormat, target)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}
	return date
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	daysUntil := time.Until(target).Hours() / 24
	return daysUntil
}

func main() {
	bday := flag.String("bday", "2023-09-03", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
