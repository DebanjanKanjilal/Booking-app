package functions

import "fmt"

// function for printing only the first names
func PrintFirstNames(bookings []UserData) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	fmt.Println("All our bookings:-", firstNames)
}
