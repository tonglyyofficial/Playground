package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) { /// Capital letter to export function from another file
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
