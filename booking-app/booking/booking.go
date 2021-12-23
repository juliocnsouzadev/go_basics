package booking

import (
	"fmt"
	"net/mail"
)

type Booking struct {
	Name     string
	LastName string
	Email    string
	Tickets  uint
}

func validateByLength(label string, name string, minLength int) bool {
	valid := len(name) >= minLength
	if !valid {
		fmt.Printf("Please inform a %v with at least %v characters long.\n", label, minLength)
	}
	return valid
}

func validateMin(label string, value uint, min uint) bool {
	valid := value >= min
	if !valid {
		fmt.Printf("Min of %v alowed to %v.\n", min, label)
	}
	return valid
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	valid := err == nil
	if !valid {
		fmt.Printf("Please enter a email address e.g. ticket@conf.com\n")
	}
	return err == nil
}

func Read(remainingTickets uint) *Booking {
	var booking = Booking{"", "", "", 0}
	fmt.Println("---------------------------------------------")

	fmt.Println("What's your name?")
	fmt.Scan(&booking.Name)
	for !validateByLength("Name", booking.Name, 3) {
		fmt.Scan(&booking.Name)
	}

	fmt.Println("What's your last name?")
	fmt.Scan(&booking.LastName)
	for !validateByLength("Last name", booking.LastName, 3) {
		fmt.Scan(&booking.LastName)
	}

	fmt.Println("What's your email?")
	fmt.Scan(&booking.Email)
	for !validEmail(booking.Email) {
		fmt.Scan(&booking.Email)
	}

	fmt.Println("How many tickets do you want to book?")
	fmt.Scan(&booking.Tickets)
	for !validateMin("Ticket", booking.Tickets, 1) {
		fmt.Scan(&booking.Tickets)
	}

	if remainingTickets < booking.Tickets {
		var decison string = ""
		for decison != "YES" && decison != "NO" {
			fmt.Printf("Sorry, only %v tickets available. Type NO to cancel or YES to accept buying the %v tickets aviable:\n", remainingTickets, remainingTickets)
			fmt.Scan(&decison)
		}
		fmt.Println("---------------------------------------------")
		if decison == "NO" {
			fmt.Printf("Canceled By User. Bye!\n")
			return nil
		} else {
			booking.Tickets = remainingTickets
		}
	} else {
		fmt.Println("---------------------------------------------")
	}
	return &booking
}
