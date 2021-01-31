package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sum(n1, n2 int) int {
	return n1 + n2
}
func main() {
	fmt.Println(greeting("World"))
	fmt.Println(sum(1, 2))
}
