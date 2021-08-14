package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a message and print it
	message := greetings.HiPerson("Gladys")
	message2 := greetings.ByePerson("Gladys")
	fmt.Println(message)
	fmt.Println(message2)
}
