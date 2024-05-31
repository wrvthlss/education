package main

import (
	"fmt"
	"math/rand"
	"time"
)

// worker function simulates processing a task and sending the result back.
func worker(id int, tasks <-chan int, results chan<- int, done chan struct{}) {
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				// Channel closed
				fmt.Printf("Worker %d shutting down\n", id)
				return
			}
			// Simulate work by sleeping for a random time
			sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
			fmt.Printf("Worker %d processing task %d (sleep %v)\n", id, task, sleepTime)
			time.Sleep(sleepTime)
			results <- task // Send result back

		case <-done:
			fmt.Printf("Worker %d received done signal\n", id)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numWorkers := 3
	numTasks := 10

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)
	done := make(chan struct{})

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results, done)
	}

	// Send tasks to the workers
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)

	// Collect results with a timeout
	timeout := time.After(3 * time.Second)
	for i := 1; i <= numTasks; i++ {
		select {
		case result := <-results:
			fmt.Printf("Received result: %d\n", result)
		case <-timeout:
			fmt.Println("Timeout reached, stopping collection of results")
			close(done)
			return
		}
	}

	// Allow workers to shut down gracefully
	close(done)
	time.Sleep(1 * time.Second) // Give workers time to print shutdown messages
	fmt.Println("All tasks processed.")
}
