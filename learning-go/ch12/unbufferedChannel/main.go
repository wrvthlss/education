package main

import (
	"fmt"
	"sync"
)

func doubler(x int) int {
	return x * 2
}

// Producer generates numbers and sends them to the channel.
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
	}
	// Close channel after sending all values.
	close(ch)
}

// Consumer receives the numbers from the channel and processes them.
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Consuming and Doubling: %d\n", doubler(val))
	}
}

func main() {

	// Create an unbuffered channel
	ch := make(chan int)

	// Create a wait group.
	var wg sync.WaitGroup

	// Add two to the WaitGroup, one for each go routine.
	wg.Add(2)

	// Start the producer and consumer goroutines.
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
