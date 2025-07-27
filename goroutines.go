package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulating a database call with goroutines
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	startTime := time.Now()
	for index := 0; index < len(dbData); index++ {
	dbCall(index)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(startTime))
}

func dbCall(index int) {
	delay := rand.Float32() * 2000 // Simulating a delay of up to 2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond) // This pauses the program for the random delay
	result := dbData[index] // Result from the "database"
	fmt.Println("The result from the database is:", result) // print call
}
