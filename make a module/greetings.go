//	Tutorial from golang.org on how to create a module for others to use.
//
//	Step 1: create a go mod init for example.com/greetings
//	Step 2: code below.

package greetings
//	Declaring a package to collect related functions.

import "fmt"

//	Hello returns a greeting for the named person.
//	A function whose name starts with a capital letter can be called
//		by a function not in the same package.
//	Known as a "exported name"
func Hello (name string) string {
	//	Return a greeting embeding the name in a message.
	//	:= operator is a shortcut for declaring an initializing
	//		a var in one line.  Otherwise below wold be:
	
	//	var message string
	//	message = ...
	
	message := fmt.Sprintf("Hi, %v.  Welcome!", name)
	//	Sprintf which has a format string, where %v (verb)
	//		 subbed for name.
	return message
}
