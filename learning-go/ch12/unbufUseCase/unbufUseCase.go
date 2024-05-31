package main

import (
	"fmt"
	"sync"
)

// worker function
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		results <- task * 2
	}

}

func main() {

	tasks := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send 5 tasks
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Collect results
	go func() {
		for i := 1; i <= 5; i++ {
			result := <-results
			fmt.Printf("Received reults %d\n", result)
		}
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(results)

}
