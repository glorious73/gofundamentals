package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50

	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of", totalTickets, "and we have", remainingTickets, "remaining.")
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var userTickets uint

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Please enter number of tickets to be booked")
		fmt.Scan(&userTickets)
		fmt.Printf("%v booked %v tickets.\n", firstName, userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Remaining tickets for conference are %v\n", remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)
			var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names are %v\n", firstNames)
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			continue
		}
	}
	fmt.Println("Tickets are now sold out. Please comeback next year.")
}
