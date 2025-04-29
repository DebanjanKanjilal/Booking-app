// function for greeting the users
package functions

import "fmt"

func GreetUser(occasionName string, numberOfTickets uint, remainingTickets int) {
	fmt.Printf("Welcome to the %v booking application!!!\n", occasionName)
	fmt.Printf("We have total of %v seats and %v tickets are still remaining.\n", numberOfTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
