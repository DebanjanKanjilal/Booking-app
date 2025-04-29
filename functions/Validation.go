// function for validating the user inputs
package functions

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTicketsNumber uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicketsNumber > 0 && userTicketsNumber <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
