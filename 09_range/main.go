package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 60, 80}
	for i, id := range ids {
		fmt.Printf("ID:%d, index %d\n", id, i)
	}
	fmt.Println()

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum ", sum)
	fmt.Println()

	phones := map[string]string{"Bob": "12345", "Sharon": "54321"}

	for key, value := range phones {
		fmt.Printf("%s:%s\n", key, value)
	}

	for name := range phones {
		fmt.Println(name)
	}
}
