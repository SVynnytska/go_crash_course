package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Sharon")
	fmt.Println(emails)

	phones := map[string]string{"Bob": "12345", "Sharon": "54321"}
	fmt.Println(phones)

}
