//	Small program for installing a package (quote)
//	When run, will downlaod package (if needed)
//		and will print out the following quote:
//	"Don't communicate by sharing memory,
//		share memory by communicating."

package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}
