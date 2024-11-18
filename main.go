package main

import "fmt"

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	// define a variable with := it will understand that its a var and get its type from the value we get to it
	remainedTicket := 50 

	var array [50]string

	var slice []string

	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >\n", ticketNumber, remainedTicket)
	var firstName string
	var lastName string
	var email string
	fmt.Println("Please Enter your firstname: ")
	// in golang you get the input with Scan function like java :) ; also you should user pointer to access to the variable
	fmt.Scan(&firstName)
	fmt.Println("Please Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("Please Enter your email: ")
	fmt.Scan(&email)
	array[0] = firstName
	slice = append(slice, firstName + " " + lastName)

	fmt.Printf("This is the array value : %v\n", array)

	fmt.Printf("This is the slice value : %v\n", slice)

	fmt.Printf("< Thanks for booking a ticket %v %v .The confirmation email will send to %v >", firstName, lastName, email)

	// you can define a variable and get a type to it and then get value to it when you want this you should define variable with type like line 13

}
