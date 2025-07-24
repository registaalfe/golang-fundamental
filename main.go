// 10.02

package main

import "fmt"

func main() {
	firstName := "John"
	lastName := "Doe"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'", firstName, lastName) // better used Printf than Println for formatted output
}