package main

import (
	"fmt"
	"leap_lib"
	"math/rand"
)

func main() {
	maxYear := 2100
	minYear := 2000
	year := rand.Intn(maxYear-minYear) + minYear
	isLeapYear := leap_lib.Leap_year(year)

	fmt.Printf("%v leap year: %v", year, isLeapYear)
}
