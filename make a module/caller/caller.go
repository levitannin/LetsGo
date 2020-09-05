//	This code will call the greetings package made earlier.
//	Updated to take an error code into a log function.

package main

//	Method for importing multiple packages.
import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	//	Set properties of predefined Logger, including
	//		the log entry prefix and flag to disable
	//		printing, time, source file, and line num
	log.StPrefix("greetings: ")
	log.SetFlags(0)


	//	Get a greeting message and print it!
	message := greetings.Hello("Levi")
	//	if error returned, print to consol and exit program.
	if err != nil{
		log.Fatal(err)
	}
	
	//	if no error returned, print returned message to console.
	fmt.Println(message)
}
