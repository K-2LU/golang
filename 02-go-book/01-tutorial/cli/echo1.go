package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	s, sep = "", ""
	// secondary for loop
	// for index(_), value(arg)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// alternate join operation
	fmt.Println(strings.Join(os.Args[1:], " "))
}
