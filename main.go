package main

// imported libraries
import (
	"booking-app/functions"
	"fmt"
)

// package level variables
const totalNumberOfTickets int = 50

var occasionName string = "GO Conference"
var remainingTickets uint = 50

var bookings = make([]functions.UserData, 0)

// main function
func main() {

	// greet user by printing welcoming message
	functions.GreetUser(occasionName, remainingTickets, totalNumberOfTickets)

	for {

		// first,last names, email, number of tickets user wants to buy
		firstName, lastName, email, userTicketsNumber := functions.GetUserInput()

		// validation boolean values
		isValidName, isValidEmail, isValidTicketNumber := functions.ValidateUserInputs(firstName, lastName, email, userTicketsNumber, remainingTickets)

		// validate both first and last names to be not too short
		if !isValidName {
			fmt.Println("Wrong name. Enter proper First and last names.")
			continue
		}

		// validate emnail if it contains "@" character
		if !isValidEmail {
			fmt.Println("Wrong email. Enter a proper EmailID.")
			continue
		}

		// validate the number that user is entering for number of tickets
		// the value shouldn't be less than 0 or more than the remaining tickets
		if isValidTicketNumber {
			// book ticket and keeping count of the remaining tickets
			bookings = functions.BookTicket(firstName, lastName, occasionName, bookings, &remainingTickets, userTicketsNumber, email)

			// send tickets after a delay of 10 seconds
			// added go in front to make this function concurrent/uninterruptible
			functions.Wg.Add(1)
			go functions.SendTicket(firstName, lastName, occasionName, userTicketsNumber, email)

			// print only first names to partially protect user identity
			functions.PrintFirstNames(bookings)

			// stop the booking process if all 50 seats are booked
			if remainingTickets == 0 {
				fmt.Println("Our conference is housefull. Please come back for the next event.")
				break
			}
		} else {
			// if they enter number of purchasing tickets more than remaining tickets
			fmt.Printf("We have only %v seats, so booking %v tickets is not possible. Try again properly.\n", remainingTickets, userTicketsNumber)
		}
	}
	functions.Wg.Wait()
}
