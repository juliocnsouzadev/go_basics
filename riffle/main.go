package main

import (
	"log"
	"time"
)

func main() {
	entries, err := importData(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
