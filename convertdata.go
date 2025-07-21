package main

import "fmt"

func main() {
	var data32 int32 = 48445 // declare a 32-bit integer
	
	// Step 1: Convert int32 to int64
	var data64 int64 = int64(data32)
	
	// Step 2: Convert int64 to int16
	var data16 int16 = int16(data64)

	// Print the values
	fmt.Println("int32:", data32)
	fmt.Println("int64:", data64)
	fmt.Println("int16:", data16)
}
