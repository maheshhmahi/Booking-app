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

	// fmt.Printf("Type of conferenceName %T\n", conferenceName)

	greeUsers(conferenceName, conferenceTickets, remainingTickets)

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := noOfTickets > 0 && noOfTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= noOfTickets

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, noOfTickets, email)

			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

			//call print first names
			firstNames := getFirstNames(bookings)
			fmt.Printf("These first names of our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is wrong")
			}
		}
	}
}

func greeUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		fName := strings.Split(booking, " ")
		firstNames = append(firstNames, fName[0])
	}

	return firstNames
}
