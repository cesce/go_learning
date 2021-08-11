package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const randomMax = 10
	var counter = 10
	var randomNumber = 0
	rand.Seed(time.Now().UnixNano())

	for counter >= 0 {
		randomNumber = rand.Intn(randomMax)
		if randomNumber == 0 {
			fmt.Printf("Count Down Aborted - Random: %v", randomNumber)
			break
		} else {
			time.Sleep(time.Second)
			fmt.Printf("Count %v - Random: %v\n", counter, randomNumber)
		}
		counter--
	}
}
