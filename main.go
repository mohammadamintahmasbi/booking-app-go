package main

import "fmt"

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	// define a variable with := it will understand that its a var and get its type from the value we get to it
	remainedTicket := 50 

	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >", ticketNumber, remainedTicket)
	var userName string

	fmt.Println(userName)
	// you can define a variable and get a type to it and then get value to it when you want this you should define variable with type like line 13
	userName = "Username"

}
