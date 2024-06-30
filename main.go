package main

import (
	"fmt"
)

const totalTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 1)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

		var userData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}

		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("%v booked %v tickets.\n", firstName, userTickets)
			fmt.Printf("Remaining tickets for conference are %v\n", remainingTickets)

			bookings = append(bookings, userData)
			fmt.Printf("Bookings are now:\n%v\n", bookings)

			var firstNames []string = getFirstNames()
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

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets to be booked")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
