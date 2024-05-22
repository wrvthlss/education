package main

import (
	"fmt"
	"sync"
)

// Task struct represents a simple task with an ID.
type Task struct {
	ID int
}

func (x *Task) dblr() Task {
	res := x.ID * 2
	return Task{
		ID: res,
	}
}

// producer generates tasks and sends them to the channel.
func producer(ch chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 10; i++ {
		task := Task{ID: i}

		fmt.Printf("Producing task: %d\n ", task.ID)

		// Write task values to channel.
		ch <- task
	}
	close(ch)
}

// consumer receives tasks from the channel and processes them.
func consumer(ch chan Task, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range ch {
		fmt.Printf("Consumer %d processing task: %d\n", id, task.dblr())
	}
}

func main() {

	// Create a buffered channel with capacity 5.
	ch := make(chan Task, 5)

	// Create a WaitGroup.
	var wg sync.WaitGroup

	// Add to the WaitGroup counter.
	wg.Add(1) // for the producer

	//  Start the producer goroutine.
	go producer(ch, &wg)

	// Start multiple consumer goroutines.
	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(ch, i, &wg)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

}
