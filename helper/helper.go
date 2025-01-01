package helper
// when you add your file to main package you can access to it without any import,
// in fact you can access to a functions in the same packages
//if you want to create new package you must create new folder and add your go file there and then import your package in the main file
import "fmt"
//super-simple function
func Greeting(conferenceName string, ticketNumber int, remainedTicket int){
	fmt.Println("Welcome to booking ", conferenceName, "app !")
	fmt.Printf("< %v ticket has planed for this conference and %v is available >\n", ticketNumber, remainedTicket)
}

// to access to a function in a package from outside you must export it but how?
// like java with its private and public and protect function's type you can export and fucntion and access to it with starting it name with capital letter
// if you change "greeting" to "Greeting" you change it to accessiable
// now you understand why the fmt or string fuction start with capital letter. yes?