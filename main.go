package main

import (
	"fmt"
	"strings"
)

func main() {
	var eventName = "Go conference"
	var eventTickets uint = 100
	greetUser(eventName, eventTickets)
	
	

	// fmt.Printf("eventName is type %T, eventTickets is type %T\n", eventName, eventTickets)
	// fmt.Printf("hello people, Only %v is available\n", eventName)
	// fmt.Printf("Get your %v tickets here only %v tickets available\n", eventName, eventTickets)

	// var userName string
	// var userTickets int
	// userName = "Adedotun"
	// userTickets=30

	// println(userName)
	// println(userTickets)

	// var userAddress int

	// fmt.Println("Enter your ticket number")
	// fmt.Scan(&userAddress)
	// fmt.Printf("Booked at %v", userAddress)

	//booking a ticket
	// var firstName string
	// var lastName string
	// var email string
	// var ticketsBooked uint
	// var availableTickets uint= 50

	// println("Input first name:")
	// fmt.Scan(&firstName)
	// println("Input last name:")
	// fmt.Scan(&lastName)
	// println("Input email:")
	// fmt.Scan(&email)
	// println("input number of tickets:")
	// fmt.Scan(&ticketsBooked)

	// fmt.Printf("Hi %v %v, thanks for buying %v tickets,a payment receipt will be sent your mail %v\n", firstName, lastName, ticketsBooked, email)

	// availableTickets = availableTickets - ticketsBooked
	// fmt.Printf("only %v tickets remaining for %v\n", availableTickets, eventName)

	//Arrays
	var bookings  []string
	// bookings = append(bookings, "Adedotun")
	// fmt.Printf("The whole array is: %v\n", bookings)
	// fmt.Printf("The array length is: %v\n", len(bookings))

	
	
	//loop
	for{
		var firstName string
		var lastName string
		var email string
		var ticketsBooked uint
		var availableTickets uint= 50
		firstnames:=returnFirstNames(bookings, firstName, lastName)
		fmt.Printf("These are the first names of booking: %v\n", firstnames)

		
		println("Input first name:")
		fmt.Scan(&firstName)
		println("Input last name:")
		fmt.Scan(&lastName)
		println("Input email:")
		fmt.Scan(&email)
		println("input number of tickets:")
		fmt.Scan(&ticketsBooked)

		isValidEmail := strings.Contains(email, "@")


		if isValidEmail {
			fmt.Printf("Hi %v %v, thanks for buying %v tickets,a payment receipt will be sent your mail %v\n", firstName, lastName, ticketsBooked, email)
	
			availableTickets = availableTickets - ticketsBooked
			fmt.Printf("only %v tickets remaining for %v\n", availableTickets, eventName)
	
			
		//IF STATEMENTS

			noAvailableTicket := availableTickets == 0
			if noAvailableTicket{
				fmt.Println("Our tickets are sold out at the moment, please check back next week")
				break
			}
		} else{
			if !isValidEmail{

				fmt.Printf("your email input is invalid\n")
			}
			continue
		}

		//SWITCH STATEMENT

		// city:= "Nigeria"
		// switch city{
		// 	case "Nigeria":
				//run code logic for Nigeria
			// case "London":
				//run code logic for london
			// default:
				//run default code logic
		// }
	}
}

func greetUser(eventName string, eventTickets uint ){
	fmt.Printf("Welcome to our %v, we have %v available tickets to book\n", eventName, eventTickets)

}

func returnFirstNames(bookings []string, firstName string, lastName string) []string{
	bookings = append(bookings, firstName + " "+ lastName)
		
	firstnames := []string{}
	
	for _, booking := range bookings{
		var names = strings.Fields(booking)
		
		firstnames = append(firstnames,names[0])
	}
	return firstnames
}