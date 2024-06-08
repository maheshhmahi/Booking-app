package main

import (
	"fmt"
)

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

	var firstName string
	var lastName string
	var email string
	var noOfTickets uint

	// ask user his name
	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets:")
	fmt.Scan(&noOfTickets)

	remainingTickets -= noOfTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, noOfTickets, email)

	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}
