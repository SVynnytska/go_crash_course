package main

import (
	"fmt"
)

// var name = "Sofiia"

func main() {

	const age = 28
	var isCool = true
	isCool = false

	//Shorthand
	// name := "Sofiia"
	// size := 5.5
	name, size := "Sofiia", 4.6

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
