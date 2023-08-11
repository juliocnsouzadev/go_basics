package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "data/riffle_entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData(path string) ([]raffleEntry, error) {

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var entries []raffleEntry
	if err := json.Unmarshal(fileBytes, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

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
