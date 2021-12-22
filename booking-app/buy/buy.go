package buy

import (
	"fmt"
	"net/mail"
)

type Buy struct {
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

func Read(remainingTickets uint) *Buy {
	var buy = Buy{"", "", "", 0}
	fmt.Println("---------------------------------------------")

	fmt.Println("What's your name?")
	fmt.Scan(&buy.Name)
	for !validateByLength("Name", buy.Name, 3) {
		fmt.Scan(&buy.Name)
	}

	fmt.Println("What's your last name?")
	fmt.Scan(&buy.LastName)
	for !validateByLength("Last name", buy.LastName, 3) {
		fmt.Scan(&buy.LastName)
	}

	fmt.Println("What's your email?")
	fmt.Scan(&buy.Email)
	for !validEmail(buy.Email) {
		fmt.Scan(&buy.Email)
	}

	fmt.Println("How many tickets do you want to book?")
	fmt.Scan(&buy.Tickets)
	for !validateMin("Ticket", buy.Tickets, 1) {
		fmt.Scan(&buy.Tickets)
	}

	if remainingTickets < buy.Tickets {
		var decison string = ""
		for decison != "YES" && decison != "NO" {
			fmt.Printf("Sorry, only %v tickets available. Type NO to cancel or YES to accept buying the %v tickets aviable:\n", remainingTickets, remainingTickets)
			fmt.Scan(&decison)
		}
		fmt.Println("---------------------------------------------")
		if decison == "NO" {
			fmt.Println("Canceled By User. Bye!\n")
			return nil
		} else {
			buy.Tickets = remainingTickets
		}
	} else {
		fmt.Println("---------------------------------------------")
	}
	return &buy
}
