package main

import "fmt"

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	// define a variable with := it will understand that its a var and get its type from the value we get to it
	remainedTicket := 50 

	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >", ticketNumber, remainedTicket)
	var firstName string
	var lastName string
	var email string

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&email)

	fmt.Printf("< Thanks for booking a ticket %v %v .The confirmation email will send to %v >", firstName, lastName, email)

	fmt.Println(firstName)
	// you can define a variable and get a type to it and then get value to it when you want this you should define variable with type like line 13
	firstName = "Username"

}
