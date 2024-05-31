package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	// Anonymous goroutine (Producer).
	wg.Add(1)
	go func() {
		defer wg.Done() // Decrements the wait counter when goroutine finishes.
		inGoroutine := 1

		// select ensures send occurs only when m
		select {
		// Attempt to send value 1 to ch1.
		// Block until a corresponding receive operation on ch1.
		case ch1 <- inGoroutine:
			fmt.Println("Sent to ch1:", inGoroutine)
		}

		// receiving from ch2
		select {
		// Receive a value from ch2
		// Block until corresponding send operation on ch2.
		case fromMain := <-ch2:
			fmt.Printf("goroutine: %v, %v\n", inGoroutine, fromMain)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		inMain := 2

		select {
		// Wait to receive a value from `ch1`
		case fromGoroutine := <-ch1:
			fmt.Printf("main received from goroutine: %v\n", fromGoroutine)

			select {
			// Attempt sending the value 2 to ch2
			// Block until corresponding receive operation on ch2
			case ch2 <- inMain:
				fmt.Println("Sent to ch2:", inMain)
			}
		}
	}()

	wg.Wait() // Wait until all goroutines finish
	fmt.Println("All tasks completed.")
}
