// a simple Go based concert booking application implemeting some basic concepts
package main

import (
	"fmt"
	"time"
)

// main function
func main() {
	//declatation of variables and conctants
	appName := "Concert Booking Application" // using : instead of var. cannot be used to declare constants
	const concertTickets int = 100
	var remainingTickets uint = 100 //unsigned int i.e cannot be negative

	//user input vars
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//using printf statement to print in a specific format. print or println can also be used
	fmt.Printf("Hello and Welcome to %v\n", appName)
	fmt.Printf("Please book your tickets! %v tickets available\n", remainingTickets)
	fmt.Printf("The time is %v\n", time.Now())

	// Ask user to input data
	fmt.Printf("Enter your First Name:\n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your Last Name:\n")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your E-mail address:\n")
	fmt.Scan(&email)
	fmt.Printf("\nEnter number of tickets to book:\n")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you  %v %v for booking %v ticket(s). You will receive a confirmation E-mail at %v\n", firstName, lastName, userTickets, email)
}
