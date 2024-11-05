package main

import "fmt"

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	var remainedTicket = 50
	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >", ticketNumber, remainedTicket)
}
