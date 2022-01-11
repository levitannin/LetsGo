// This is a comment, yay!

// Packages in this code.
package main

// Imported libraries used in this code.
import (
	"fmt"
  "time"
)

// Main function of the code
func main() {
	fmt.Println("Hello, World!")
  fmt.Println("The current time is:", time.Now())

  /*
  These are functions from other go files.
  In order for these to run, the terminal call needs to change from run main.go, to include all other files.

  Example:
    go run main.go maths.go vartypes.go

    or

    go run *.go

  These are to run in commandline and mean nothing regarding building or libraries.
  */

  // maths
  random()
  maths()

  fmt.Println("If I add 13 and 849, I'll get:", addArgs(13, 849))

  fmt.Println("If I subtract 849 from 13, I'll get:", subArgs(13, 849))

  fmt.Println(split(42))

  // vartypes
  a, b := swap("This is a string.", "No this is a string.")
  fmt.Println(a,b)
  
  variables()
  varvars()
  basetype()
  zerovals()
}

/*
This is a comment block.  I reviewed
  how to make comments using this link:
    https://www.digitalocean.com/community/tutorials/how-to-write-comments-in-go
  
This series of go files are going to be my notes, so expect comments ;)
*/