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

	//using printf statement to print in a specific format. print or println can also be used
	fmt.Printf("Hello and Welcome to %v\n", appName)
	fmt.Printf("Please book your tickets! %v tickets available\n", remainingTickets)
	fmt.Printf("The time is %v\n", time.Now())

	var userName string
	var userTickets int

	// user input
	userName = "Pat"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}
