package main

import (
	"fmt"
	"strings"
)

const totalTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

func main() {

	greetUser()

	var bookings []string

	for remainingTickets > 0 {

		firstName, lastName, userTickets := getUserInput()
		isValidName, isValidTickets := validateInput(firstName, lastName, userTickets, remainingTickets)

		if isValidName && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("%v booked %v tickets.\n", firstName, userTickets)
			fmt.Printf("Remaining tickets for conference are %v\n", remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)
			var firstNames []string = getFirstNames(bookings)
			fmt.Printf("First names are %v\n", firstNames)
		} else {
			fmt.Printf("First name and last name each contain at least 2 characters and the remaining # of tickets is %v\n", remainingTickets)
			continue
		}
	}
	fmt.Println("Tickets are now sold out. Please comeback next year.")
}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of", totalTickets, "and we have", remainingTickets, "remaining.")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	var firstNames []string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter number of tickets to be booked")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func validateInput(firstName string, lastName string, userTickets uint, remainingTickets uint) (bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidTickets bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTickets
}
