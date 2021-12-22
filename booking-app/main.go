package main

import (
	"booking-app/buy"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	clear()

	const conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	var bookings []*buy.Buy

	for remainingTickets > 0 {
		fmt.Printf("\nWelcome to %v application\n", conferenceName)
		fmt.Printf("Get your ticket today! Only %v tickets left!\n", remainingTickets)
		var buy = buy.Read(remainingTickets)
		if buy != nil {
			fmt.Printf("%v %v bought %v tickets\nWe will send the tickets to %v. Thank you!\n", buy.Name, buy.LastName, buy.Tickets, buy.Email)
			remainingTickets = remainingTickets - buy.Tickets
			bookings = append(bookings, buy)
		}
		wait("... getting ready for next booking in... ", 5)
		clear()
	}

	fmt.Printf("\n\n*** All Tickets are Sold out!***\n")
	wait("... generating bookings report in... ", 5)
	clear()
	fmt.Printf("\n*** Total of %v Bookings ***\n", len(bookings))
	for _, book := range bookings {
		fmt.Printf("> %03d tickets bought by %v %v \n", book.Tickets, book.Name, book.LastName)
	}
}

func wait(prefix string, seconds int) {
	counter := seconds
	fmt.Printf(prefix)
	for counter >= 0 {
		fmt.Printf("%v... ", counter)
		duration := time.Second
		time.Sleep(duration)
		counter--
	}
	fmt.Printf("\n")
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
