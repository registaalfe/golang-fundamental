package main

import "fmt"

func main() {
	var c = make(chan int) // Create a channel to communicate between goroutines with data type int
	go process(c) // Keyword 'go' is running on the background
	for i := range c { // Loop is for receiving data from the channel
		fmt.Println(i)
	}
}

func process(c chan int) { // Receive parameter 'c' of type channel int
	defer close(c) //Schedules the channel to be closed when the function returns
	for i := 0; i < 5; i++ {
		c <- i // Send data to the channel
	}
}
