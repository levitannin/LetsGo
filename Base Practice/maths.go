// Packages used in this code.
package main

// Imported paths/libraries
import (
  "fmt"
  "math"
  "math/rand"
  "time"
)

func random(){
  fmt.Println("This system's favorite number is:", rand.Intn(100))

  //The way random is called can also be seeded.
  rand.Seed(13)
  fmt.Println("This is a seeded random number:", rand.Intn(100))

  //This is a random number using current system time as the seed
  rand.Seed(time.Now().UnixNano())
  fmt.Println("These are seeded random numbers:", rand.Intn(100), rand.Intn(100), rand.Intn(100))
}

/*
Interesting tutorial: 
  https://www.includehelp.com/golang/demonstrate-the-rand-seed-function.aspx
*/

/*
Exports begin with capital letters.  Example:
  Pi is an exported name of the math package.

Example below will throw an error for lowercase vs uppercase of the export name.
*/

func maths(){
  //fmt.Println("This will fail:", math.pi)
  fmt.Println("This will not fail:", math.Pi)
}

/*
This function focuses on how to work with arguments. Since it works with maths, it gets to hang out here.
*/

func addArgs(x int, y int) int{
  return x + y
}

// This function shows another syntax for arguments.

func subArgs(x, y int) int{
  return x - y
}

/*
Previously our functions returned nothing or returned a single int.  This time we are returning two int so we need parantheses around (x,y int) similar to when we declare these variables as arguments.
*/
func split(sum int)(x, y int){
  x = sum * 4 / 9
  y = sum - x
  
  // to return the new variables add a return call.
  return
}