package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int) // Create a channel with data type int
	go sendData([]int{1, 2, 3, 4}, intChan) // Start a goroutine to send data to the channel with data static type int
	for v := range intChan {
		fmt.Println("Angka:", v)
	}

	strChan := make(chan string) // Create a channel with data type string
	go sendData([]string{"halo", "dunia"}, strChan) // Start a goroutine to send data to the channel with data static type string
	for v := range strChan {
		fmt.Println("Teks:", v)
	}
}

func sendData[T any](slice []T, ch chan T) {
	for _, v := range slice {
		ch <- v
	}
	close(ch)
}