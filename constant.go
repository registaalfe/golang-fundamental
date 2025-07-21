package main

import "fmt"

func main() {
	const nameMenu string = "Ogliolio Mie Star"
	const categoryMenu = "Spaghetti"

	// Defining multiple constants
	// const (
	// 	firstName = "Cant changed"
	// 	lastName  = "Cant changed"
	// )

	fmt.Println("What menu?", nameMenu)
	fmt.Println("This is the category:", categoryMenu)
}