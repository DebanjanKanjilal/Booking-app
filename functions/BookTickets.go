package functions

import (
	"fmt"
)

// Ticket booking function
func BookTicket(firstName string, lastName string, occasionName string, bookings []UserData, remainingTickets *uint, userTicketsNumber uint, email string) []UserData {
	// keeping count of the remaining tickets and maintaining a list of ticket holders
	*remainingTickets -= userTicketsNumber

	// create object for UserData
	var userData = UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTicketsNumber,
	}

	// create a map for user details
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailID"] = email
	// userData["numberOfTicketsThisUserPurchased"] = strconv.FormatUint(uint64(userTicketsNumber), 10)

	// add userData of every iteration in the bookings struct list
	bookings = append(bookings, userData)

	// fmt.Printf("List of Attendees:%v", bookings)

	// print thanks for successfull ticket booking
	fmt.Printf("Thank you %v %v for successfully purchasing %v ticket(s) for the %v! You will receive a confirmation mail and details of the event at %v.\n", firstName, lastName, userTicketsNumber, occasionName, email)
	// and print number of remaining tickets
	fmt.Printf("%v tickets remaining for %v.\n", *remainingTickets, occasionName)
	return bookings
}
