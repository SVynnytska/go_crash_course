package main

import "fmt"

func main() {
	x := 6
	y := 8
	if x <= y {
		fmt.Printf("%d is less then or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d\n", y, x)
	}

	color := "red"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}
