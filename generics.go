package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3} // Data static type int
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3} // Data static type float32
	fmt.Println(sumSlice[float32](float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Ini kalo gapake generics
// func sumIntSlice(slice []int) int {
// 	var sum int
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }

// func sumFloat32Slice(slice []float32) float32 {
// 	var sum float32
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }
