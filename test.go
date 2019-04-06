package main

import (
	"fmt"
	"github.com/AzharSindhi/golangexamples"
	)

func main() {
	slice := [] byte{'h','e','l','l','o'}
	
	encrypted := golangexamples.Encrypt(slice, 3)
	concatenated := golangexamples.ConcatSlice(slice)
	name := "Azhar"
	greetings := golangexamples.Greetings(name)
	fmt.Println(encrypted)
	fmt.Println(concatenated)
	fmt.Println(greetings)
}
