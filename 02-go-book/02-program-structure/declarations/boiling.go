package main

import "fmt"

// package level declartion
// visible to all the files of the
// package
const boilingF = 212.0

func main() {
	// local to main func
	// only visible to the func
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)

	ftoc()
}
