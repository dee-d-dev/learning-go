package main

import "fmt"

func main() {
	var eventName = "Go conference"
	// var eventTickets uint = 100
	

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
	var firstName string
	var lastName string
	var email string
	var ticketsBooked uint
	var availableTickets uint= 50

	println("Input first name:")
	fmt.Scan(&firstName)
	println("Input last name:")
	fmt.Scan(&lastName)
	println("Input email:")
	fmt.Scan(&email)
	println("input number of tickets:")
	fmt.Scan(&ticketsBooked)

	fmt.Printf("Hi %v %v, thanks for buying %v tickets, payment receipt will be sent your mail %v\n", firstName, lastName, ticketsBooked, email)

	availableTickets = availableTickets - ticketsBooked
	fmt.Printf("only %v tickets remaining for %v", availableTickets, eventName)


}