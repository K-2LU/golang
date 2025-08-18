package main

import "fmt"

func main() {
	var intArr [3]int
	fmt.Println(intArr[2])
	fmt.Println(intArr[0:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int = [3]int{1, 2, 3}
	// intArr2  := [...]int{1,2,3}	// also valid
	fmt.Println(intArr2)
}
