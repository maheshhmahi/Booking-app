package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = []string{}

// creating list of bookings of map. we need to initialize the size first
var bookings = make([]map[string]string, 0)

func main() {

	//var conferenceName = "Go Conference"
	// also can be declared as below

	// defining the array
	// var bookings = [50]string{"Mahesh", "Shruthi"}
	// var bookings [50]string
	// var bookings [conferenceTickets]string

	//slice: is like dynamic array/arrayList
	// var bookings []string
	// var bookings = []string{}

	// fmt.Printf("Type of conferenceName %T\n", conferenceName)

	greetUsers()

	for {
		firstName, lastName, email, noOfTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, noOfTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, noOfTickets, firstName, lastName, email)

			//call print first names
			firstNames := getFirstNames()
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
			fmt.Println("Please try again")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, noOfTickets
}

func bookTicket(remainingTickets uint, noOfTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= noOfTickets

	//create user map
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	user["numOfTickets"] = strconv.FormatUint(uint64(noOfTickets), 10)

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, noOfTickets, email)

	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}
