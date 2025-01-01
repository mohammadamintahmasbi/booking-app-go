package main

import "fmt"
//super-simple function
func greeting(conferenceName string, ticketNumber int, remainedTicket int){
	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >\n", ticketNumber, remainedTicket)
}