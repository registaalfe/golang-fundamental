package main

import "fmt"

type Address struct {
	City string
}

func main() {
	address1 := Address{"Jatim"}
	address2 := address1 // If u don't use the address operator (&), address2 will not be a copy of address1

	address2.City = "Bali"

	fmt.Println(address1) // Output: Bali
	fmt.Println(address2) // Output: Bali
}