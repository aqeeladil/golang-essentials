package main

import (
	"fmt"
	"sync"
	"time"
)

// Mock function to simulate a task (e.g., downloading a file)
func downloadFile(id int) {
	fmt.Printf("Task %d: Downloading file...\n", id)
	time.Sleep(2 * time.Second) // Simulate 2-second task
	fmt.Printf("Task %d: Download complete!\n", id)
}

func main() {
	// Measure execution time
	start := time.Now()

	// Number of tasks to execute
	numTasks := 5

	// Using a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	fmt.Println("Starting tasks concurrently using goroutines...")

	// Launch tasks as goroutines
	for i := 1; i <= numTasks; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go func(id int) {
			defer wg.Done() // Decrement counter when task is done
			downloadFile(id)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Calculate total execution time
	elapsed := time.Since(start)
	fmt.Printf("All tasks completed in %s\n", elapsed)
}
