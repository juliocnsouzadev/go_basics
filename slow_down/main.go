package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	for _, token := range strings.Split(msg, " ") {
		slowedToken := ""
		for i, c := range token {
			slowedToken += strings.Repeat(string(c), i+1)
		}
		print(slowedToken)
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
