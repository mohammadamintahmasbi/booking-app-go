package main
// when you add your file to main package you can access to it without any import,
// in fact you can access to a functions in the same packages
//if you want to create new package you must create new folder and add your go file there and then import your package in the main file
import "fmt"
//super-simple function
func greeting(conferenceName string, ticketNumber int, remainedTicket int){
	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >\n", ticketNumber, remainedTicket)
}