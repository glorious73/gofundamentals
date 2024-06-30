package main

import "strings"

func ValidateInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTickets bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
