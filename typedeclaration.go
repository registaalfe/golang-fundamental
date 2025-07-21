package main

import "fmt"

func main() {
	type NoRoom string //declaring a new type NoRoom based on string
	
	var yourRoom NoRoom = "1123" // create a variable of type NoRoom
	fmt.Println("Your room number is:", yourRoom) // prints: Your room number is: 1123
	fmt.Println(NoRoom("1126")) // create a new data from old data type NoRoom
}