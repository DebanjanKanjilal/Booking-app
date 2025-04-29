package functions

import "sync"

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var Wg = sync.WaitGroup{}
