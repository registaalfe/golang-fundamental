package main

import "fmt"

func main() {
	const pi string = "Hello, World!"
	const e = 2.718

	const (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Println("Value of Pi:", pi)
	fmt.Println("Value of E:", e)

	// Uncommenting the following lines will cause an error
	// pi = 3.14159 // cannot assign to pi
	// e = 2.71828 // cannot assign to e
}