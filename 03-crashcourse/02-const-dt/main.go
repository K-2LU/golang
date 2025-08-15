package main

import (
	"errors"
	"fmt"
)

const piValue float32 = 3.1416

func areaCircle(radius int) float32 {
	var r float32 = float32(radius)
	return r * r * piValue
}

func intDiv(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var res int = numerator / denominator
	var rem int = numerator % denominator

	return res, rem, err
}

func main() {
	var intNum int = 1923
	// integer type: 16, 32, 64 bit
	// default: 32 bit
	fmt.Println(intNum)

	var int16Num int16 = 32767
	int16Num += 1 // will produce negative
	fmt.Printf("%d\n", int16Num)

	var floatNum float64 = 12344567.090
	fmt.Printf("%f\n", floatNum)

	var sum float64 = floatNum + float64(intNum)
	// must typecast like c, c++ or any
	// language that makes sense
	fmt.Printf("%f\n", sum)

	var1, var2 := 1, 2 // also valid
	// use when it makes intuitively
	// add type when not, good practice
	fmt.Printf("%d\t%d\n", var1, var2)

	const myConst string = "Change it to see and error"
	// myConst = "I dare you"

	fmt.Printf("%s\n", myConst)

	var area float32 = areaCircle(10)
	fmt.Printf("Area: %f\n", area)

	var res, rem, err = intDiv(15, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else if rem == 0 {
		fmt.Printf("res: %d\n", res)
	} else {
		fmt.Printf("res: %d\nrem: %d\n", res, rem)
	}

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("res: %d\n", res)
	default:
		fmt.Printf("res: %d\nrem: %d\n", res, rem)
	}

	switch rem {
	case 0:
		fmt.Println("exact division")
	case 1, 2:
		fmt.Println("close division")
	default:
		fmt.Println("not close")
	}
}
