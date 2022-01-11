// Packages used in this code.
package main

// Imported paths/libraries
import (
  "fmt"
  "math/cmplx"
)

var i, j int = 1, 2

var (
  ToBe  bool  = false
  MaxInt  uint64  =1<<64 - 1
  z complex128  =cmplx.Sqrt(-5 + 12i)
)

// This function returns two strings rather than one return object which would not require (return).
func swap (x, y string) (string, string){
  return y, x
}

func variables(){
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}

func varvars(){
// Different ways to write variables
var m, n int = 3, 6
// This short hand only works inside functions.
k := 3
c, python, java := true, false, "no!"

fmt.Println(m, n, k, c, python, java)
}

func basetype(){
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
Go's basic types are:

  bool

  string

  int  int8  int16  int32  int64
  uint uint8 uint16 uint32 uint64 uintptr

  byte // alias for uint8

  rune // alias for int32
     // represents a Unicode code point

  float32 float64

  complex64 complex128
*/

func zerovals(){
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}