package main

import (
	"fmt"
	"strings"
	"strconv"
	"booking-app/helper" // when you create a new package you can add it like this with the module name;
)

func main(){
	var conferenceName = "GO-LANG Conference"
	const ticketNumber = 50
	// define a variable with := it will understand that its a var and get its type from the value we get to it
	remainedTicket := 50
	//define the array with fix length
	var array [50]string
	// define a slice; slice is like a list in python in some parts :o
	var slice []map[string]string
	
	// there is an alternative syntax for slice definition : slice := []string

	var firstName string
	var lastName string
	var email string
	var number_of_booking_ticket int


	// in golang you get the input with Scan function like java :) ; also you should user pointer to access to the variable
	// fmt.Println("Please Enter your firstname: ")
	// fmt.Scan(&firstName)
	// fmt.Println("Please Enter your lastname: ")
	// fmt.Scan(&lastName)
	// fmt.Println("Please Enter your email: ")
	// fmt.Scan(&email)
	// you give a value to array like this
	array[0] = firstName
	// this is an append function like in python , ofcourse it's not in a class and not handel with "." but it do the same thing
	// its interesting that the go compiler recommend this syntax to you
	// slice = append(slice, firstName + " " + lastName)

	// different types of loop syntax in go

	for remainedTicket != 0 {
		//import from helper package :)
		helper.Greeting(conferenceName, ticketNumber, remainedTicket)
		firstName, lastName, email, number_of_booking_ticket = getUserInput(firstName, lastName, email, number_of_booking_ticket)
		// for assigning the output of a function if the variable did not define you should assign and define at the same time with :=
		// if you defined it before so you can assign it with "="
		isValidName, isValidEmail, isValidTicketNumber := inputValidation(
			firstName,
			lastName,
			email,
			number_of_booking_ticket,
			remainedTicket)

		if isValidName && isValidEmail && isValidTicketNumber{

			if number_of_booking_ticket <= remainedTicket{
				var userData = make(map[string]string)
				userData["firstname"] = firstName
				userData["lastname"] = lastName
				userData["email"] = email
				userData["number_of_booking_ticket"] = strconv.Itoa(number_of_booking_ticket)
				slice = append(slice, userData)
				remainedTicket = remainedTicket - number_of_booking_ticket
			}else {
				fmt.Println("You want to book tickets more than we have")
			}
		}else{
			if !isValidName{
				fmt.Println("< Your firstname and lastname must be more than 2 charecters >")
			}
			if !isValidEmail{
				fmt.Println("< Your email must contains @ sign >")
			}
			if !isValidTicketNumber{
				fmt.Println("< Your Enter invalid ticket number >")
			}
		}

	}
	fmt.Println("----------------------------------------------------------------------")
	bookingResult(array, slice, firstName, lastName, email)
	// you can define a variable and get a type to it and then get value to it when you want this you should define variable with type like line 13

}

//super-simple function with multi-return value
func getUserInput(firstName string, lastName string, email string, number_of_booking_ticket int)(string, string, string, int){
	// in golang you get the input with Scan function like java :) ; also you should user pointer to access to the variable
	fmt.Println("Please Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("Please Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("Please Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How much ticket do you want ?")
	fmt.Scan(&number_of_booking_ticket)
	
	return firstName, lastName, email, number_of_booking_ticket
}

func inputValidation(
	firstName string,
	lastName string,
	email string,
	number_of_booking_ticket int,
	remainedTicket int)(bool, bool, bool){

	isValidName := len(firstName) > 2 && len(lastName) >2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := number_of_booking_ticket > 0 && remainedTicket != 0
	
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookingResult(array [50]string, slice []map[string]string, firstName string, lastName string, email string){
	fmt.Printf("This is the booking value : %v\n", array) // array is empty. I didnot assign any thing to it.

	fmt.Printf("This is the slice value : %v\n", slice)

}