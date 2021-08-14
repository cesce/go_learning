package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a message and print it
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
