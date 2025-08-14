package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func ex() {
	// Exercise 1.1: Print the command name
	fmt.Println("Command name:", os.Args[0])

	// Exercise 1.2: Print index and value for each argument
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	// Exercise 1.3: Compare concatenation vs strings.Join
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // inefficient: creates new string each time
		sep = " "
	}
	elapsedConcat := time.Since(start)

	start = time.Now()
	joined := strings.Join(os.Args[1:], " ") // efficient
	elapsedJoin := time.Since(start)

	fmt.Println("\nConcatenated:", s)
	fmt.Println("Joined:      ", joined)
	fmt.Println("Concat time:", elapsedConcat)
	fmt.Println("Join time:  ", elapsedJoin)
}
