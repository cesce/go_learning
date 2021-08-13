package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("03_package_import")
	fmt.Println(quote.Go())
}

/*
'go run .' will produce:

main.go:6:2: no required module provides package rsc.io/quote; to add it:
        go get rsc.io/quote

We will need to run 'go mod tidy', go will add needed dependencies and download the packages
*/
