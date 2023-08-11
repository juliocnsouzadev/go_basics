package main

import (
	"flag"
	"log"
)

func main() {
	bday := flag.String("bday", "2023-09-03", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target, err := parseTime(*bday)
	if err != nil {
		log.Fatalf("Could not parse date: %s", err)
	}
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
