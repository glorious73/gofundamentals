package helper

func ValidateInput(firstName string, lastName string, userTickets uint, remainingTickets uint) (bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidTickets bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTickets
}
