package main

import (
	"errors"
	"fmt"
)

// Creating a custom error
var (
	ErrNotFound = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
)

// GetById simulates a function that retrieves an item by ID
func GetById(id string) error {
	if id == "" {
		return ErrNotFound
	}

	if id != "Go" {
		return ErrInvalidInput
	}

	return nil
}

// main function to demonstrate error handling
func main() {
	err := GetById("Go")
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Error: Item not found")
		} else if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Error: Invalid input")
		} else {
			fmt.Println("Unknown error occurred")
		}
	}
}