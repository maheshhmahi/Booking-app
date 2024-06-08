package main

import "fmt"

func main() {

	//var conferenceName = "Go Conference"
	// also can be declared as below
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Printf("Type of conferenceName %T\n", conferenceName)

	var userName string
	var userTickets int

	// ask user his name
	userName = "Mahesh"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
