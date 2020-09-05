//	This code will call the greetings package made earlier.

package main

//	Method for importing multiple packages.
import (
	"fmt"
	"example.com/greetings"
)

func main() {
	//	Get a greeting message and print it!
	message := greetings.Hello("Levi")
	fmt.Println(message)
}
