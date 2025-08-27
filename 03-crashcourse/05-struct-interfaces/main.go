package main

import "fmt"

type gasEngine struct {
	kpl    uint8
	liters uint8
	owner
}

type electricEngine struct {
	kmkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e electricEngine) kmLeft() uint8 {
	return e.kmkwh * e.kwh
}

func (e gasEngine) kmLeft() uint8 {
	return e.kpl * e.liters
}

type engine interface {
	kmLeft() uint8
}

func main() {
	var myEngine gasEngine = gasEngine{
		kpl:    50,
		liters: 12,
		owner:  owner{name: "John"},
	}

	// anonymous structs - not re-usable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 25}

	fmt.Println(myEngine.name, myEngine.kpl, myEngine.liters)
	fmt.Printf("%v kms left\n", myEngine.kmLeft())
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// fmt.Println(myEngine2)
}
