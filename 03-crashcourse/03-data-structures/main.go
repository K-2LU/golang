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

	var intSlice []int32
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{1, 2, 3}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// var instSlice3 []int32 = make(int32[], 3, 8)
	// size, capacity

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

}
