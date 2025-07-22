package main

import "fmt"

func main() {
	// Creating a map
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // Breaks out of the outer loop
	// 	}
	// 	fmt.Println("Outer loop iteration:", i)
	// }

	// Continue statement example
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Skips the rest of the loop when i is 5
		}
		fmt.Println("Current value of i:", i)
	}
}