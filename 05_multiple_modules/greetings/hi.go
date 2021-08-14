package greetings

import (
	"fmt"
)

func HiPerson(name string) string {
	message := fmt.Sprintf("Hi %v!", name)
	return message
}
