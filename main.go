package main

import "fmt"

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	// define a variable with := it will understand that its a var and get its type from the value we get to it
	remainedTicket := 50 
	//define the array with fix length
	var array [50]string
	// define a slice; slice is like a list in python in some parts :o
	var slice []string
	// there is an alternative syntax for slice definition : slice := []string

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
	// you give a value to array like this
	array[0] = firstName
	// this is an append function like in python , ofcourse it's not in a class and not handel with "." but it do the same thing
	// its interesting that the go compiler recommend this syntax to you
	slice = append(slice, firstName + " " + lastName)
	// different types of loop syntax in go
	for i := 0; i<= 5; i++{
	// in golang you get the input with Scan function like java :) ; also you should user pointer to access to the variable
		fmt.Println("Please Enter your firstname: ")
		fmt.Scan(&firstName)
		fmt.Println("Please Enter your lastname: ")
		fmt.Scan(&lastName)
		fmt.Println("Please Enter your email: ")
		fmt.Scan(&email)
		slice = append(slice, firstName + " " + lastName)
	}
	x := 2
	for x < 100{
		x = x*2
		fmt.Println(x)
	}	

	fmt.Printf("This is the array value : %v\n", array)

	fmt.Printf("This is the slice value : %v\n", slice)

	fmt.Printf("< Thanks for booking a ticket %v %v .The confirmation email will send to %v >", firstName, lastName, email)

	// you can define a variable and get a type to it and then get value to it when you want this you should define variable with type like line 13

}
