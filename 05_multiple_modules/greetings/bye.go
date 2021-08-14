package greetings

import (
	"fmt"
)

func ByePerson(name string) string {
	message := fmt.Sprintf("Bye %v!", name)
	return message
}
