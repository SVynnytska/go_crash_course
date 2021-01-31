package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "apple"
	fruits[1] = "mango"
	fmt.Println(fruits)

	veggies := [2]string{"cucumber", "popato"}
	fmt.Println(veggies)

	berries := []string{"blueberry", "raspberry", "goji"}
	fmt.Println(berries)
	fmt.Println(len(berries))
	fmt.Println(berries[1:3])
}
