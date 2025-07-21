package main

import "fmt"

func main() {
	var data = [...]int{
		88,
		89,
		90,
	}
	fmt.Println("Array contents:", data)
	fmt.Println("Length of array:", len(data))
	data[0] = 100
	fmt.Println("Updated array contents:", data)
}