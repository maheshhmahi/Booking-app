package main

import "fmt"

func main() {
	var confernceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", confernceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
