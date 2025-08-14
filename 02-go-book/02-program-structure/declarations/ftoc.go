package main

import "fmt"

func ftoc() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, converted(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, converted(boilingF))
}

func converted(f float64) float64 {
	return (f - 32) * 5 / 9
}
