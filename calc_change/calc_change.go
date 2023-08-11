package main

import (
	"log"
)

// coin contains the name and value of a coin
type coin struct {
	name  string
	value float64
}

const (
	OnePoundStr    = "1 pound"
	FiftyPenceStr  = "50 pence"
	TwentyPenceStr = "20 pence"
	TenPenceStr    = "10 pence"
	FivePenceStr   = "5 pence"
	OnePennyStr    = "1 penny"
)

var (
	OnePound    = coin{name: OnePoundStr, value: 1}
	FiftyPence  = coin{name: FiftyPenceStr, value: 0.50}
	TwentyPence = coin{name: TwentyPenceStr, value: 0.20}
	TenPence    = coin{name: TenPenceStr, value: 0.10}
	FivePence   = coin{name: FivePenceStr, value: 0.05}
	OnePenny    = coin{name: OnePennyStr, value: 0.01}

	coins = []coin{
		OnePound,
		FiftyPence,
		TwentyPence,
		TenPence,
		FivePence,
		OnePenny,
	}
)

// calculateChange returns the coins required to calculate the
func calculateChange(amount float64) map[coin]int {
	//converting to int to avoid floating point errors
	amountLeft := int(amount * 100)
	change := make(map[coin]int)
	for _, coin := range coins {

		coinValue := int(coin.value * 100)
		numberOfCoins := amountLeft / coinValue

		if numberOfCoins >= 1 {
			totalInCurrentCoin := numberOfCoins * coinValue

			amountLeft -= totalInCurrentCoin
			change[coin] = numberOfCoins

		}
		if amountLeft == 0 {
			break
		}
	}
	return change
}

// printCoins prints all the coins in the slice to the terminal.
func printCoins(change map[coin]int) {
	if len(change) == 0 {
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.")
	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}
