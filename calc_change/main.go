package main

import "flag"

func main() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for")
	flag.Parse()
	change := calculateChange(*amount)
	printCoins(change)
}
