package functions

import "fmt"

// function for getting user inputs for first and last names, email, user's number of tickets
func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	//user input of first and last names
	fmt.Println("Hi new user, give your first name:-")
	fmt.Scan(&firstName)
	fmt.Println("Give your last name:-")
	fmt.Scan(&lastName)
	// take user input for email address
	// validate emnail if it contains "@" character
	var email string
	fmt.Println("Enter your email ID:-")
	fmt.Scan(&email)
	// take user input for number of tickets user wants to purchase
	var userTicketsNumber uint
	fmt.Println("Enter the number of seats you want to book:-")
	fmt.Scan(&userTicketsNumber)
	return firstName, lastName, email, userTicketsNumber
}
