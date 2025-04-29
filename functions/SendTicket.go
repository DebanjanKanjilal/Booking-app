package functions

import (
	"fmt"
	"time"
)

func SendTicket(firstName string, lastName string, occasionName string, userTicketsNumber uint, email string) {
	// add delay for the following code to be executed after the booking
	time.Sleep(10 * time.Second)

	// content of the ticket
	var ticket = fmt.Sprintf("%v ticket(s) of %v %v for the event of %v", userTicketsNumber, firstName, lastName, occasionName)

	// notifying sending of ticket
	fmt.Println("------------------------")
	fmt.Printf("Ticket sent:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("------------------------")
	Wg.Done()
}
