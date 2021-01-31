package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstname string
	lastname  string
	city      string
	age       int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstname + " " + p.lastname + " and I'm " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person := Person{firstname: "Samantha", lastname: "Smith", city: "Boston", age: 25}

	fmt.Println(person, person)
	fmt.Println(person.greet())
	person.hasBirthday()
	fmt.Println(person.greet())
}
