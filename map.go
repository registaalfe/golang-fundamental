package main

import "fmt"

func main() {
	// Creating a map
	person := map[string]string{
		"name":    "John Doe",
		"address": "123 Main Street",
	}

	fmt.Println(person)          // Accessing value by key
	fmt.Println(person["name"])          // Accessing value by key
	fmt.Println(person["address"])          // Accessing value by key
}