package main

import (
	"booking-app/booking"
	"booking-app/conference"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const conference = conference.Conference{
	"Go Conference",
	50,
	make([]*booking.Booking)
}

const conferenceName string = "Go Conference"
const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings []*booking.Booking

func main() {
	for remainingTickets > 0 {
		greetings()
		book()
	}
	allTickesSoldOut()
	genReport()
}

func greetings() {
	clear()
	fmt.Printf("\nWelcome to %v application\n", conferenceName)
	fmt.Printf("Get your ticket today! Only %v tickets left!\n", remainingTickets)
}

func book() {
	var bookingTransaction = booking.Read(remainingTickets)
	if bookingTransaction != nil {
		fmt.Printf("%v %v bought %v tickets\nWe will send the tickets to %v. Thank you!\n", bookingTransaction.Name, bookingTransaction.LastName, bookingTransaction.Tickets, bookingTransaction.Email)
		remainingTickets = remainingTickets - bookingTransaction.Tickets
		bookings = append(bookings, bookingTransaction)
	}
	wait("... getting ready for next booking in... ", 5)
	clear()
}

func allTickesSoldOut() {
	fmt.Printf("\n\n*** All Tickets are Sold out!***\n")
	wait("... generating bookings report in... ", 5)
	clear()
}

func genReport() {
	fmt.Printf("\n*** Total of %v Bookings ***\n", len(bookings))
	for _, booking := range bookings {
		fmt.Printf("> %03d tickets bought by %v %v \n", booking.Tickets, booking.Name, booking.LastName)
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
