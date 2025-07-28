package main

import (
	"fmt"
	"sync"
	"time"
)


var m = sync.RWMutex{} // Creating a mutex to protect shared data

// Creating a WaitGroup to wait for all goroutines to finish
var wg = sync.WaitGroup{}

// Simulating a database call with goroutines
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// Creating a new slice to store results
var results = []string{}

func main() {
	startTime := time.Now()
	for index := 0; index < len(dbData); index++ {
		wg.Add(1) // Increment the WaitGroup counter
		go dbCall(index)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("\nTotal execution time: %v", time.Since(startTime))
	fmt.Printf("\nTotal number of results:", results)
}

func dbCall(index int) {
	var delay float32 = 2000 // Simulating a delay of up to 2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond) // This pauses the program for the random delay
	save(dbData[index]) // Call the save function to append the result
	log() // Call the log function to print the results
	wg.Done() // Decrement the WaitGroup counter when done
}

func save(result string) {
	m.Lock()
	results = append(results, result) // Append the result to the results slice
	m.Unlock() // Unlock the mutex after appending
}

func log() {
	m.RLock() // Lock for reading
	fmt.Println("The result from the database is:", results) // print call
	m.RUnlock() // Lock for reading
}