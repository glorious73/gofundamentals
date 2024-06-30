package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50

	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of", totalTickets, "and we have", remainingTickets, "remaining.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter number of tickets to be booked")
	fmt.Scan(&userTickets)
	fmt.Printf("%v booked %v tickets.\n", firstName, userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Remaining tickets for conference are %v\n", remainingTickets)

	bookings = append(bookings, firstName)
	fmt.Printf("All bookings: %v\n", bookings)
}
