package main

import (
	"fmt"
	"strings"
)

func main() {

	//var conferenceName = "Go Conference"
	// also can be declared as below
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// defining the array
	// var bookings = [50]string{"Mahesh", "Shruthi"}
	// var bookings [50]string
	// var bookings [conferenceTickets]string

	//slice: is like dynamic array/arrayList
	// var bookings []string
	// var bookings = []string{}
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// fmt.Printf("Type of conferenceName %T\n", conferenceName)

	for {
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

		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email, "@")
		// isValidTicketNumber := noOfTickets > 0 && noOfTickets <= remainingTickets

		if noOfTickets <= remainingTickets {
			remainingTickets -= noOfTickets

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, noOfTickets, email)

			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				fName := strings.Split(booking, " ")
				firstNames = append(firstNames, fName[0])
			}

			fmt.Printf("These first names of our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, noOfTickets)
		}

	}

}
