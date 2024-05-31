package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1

		// Place value of inGoroutine in channel.
		ch1 <- inGoroutine

		// Read value from ch2 and assign it to fromMain.
		fromMain := <-ch2
		fmt.Printf("goroutine: %v, %v\n", inGoroutine, fromMain)
	}()

	inMain := 2

	/** Deadlock fix **/
	// Receive value from ch1 before sending to ch2
	fromGoroutine := <-ch1
	fmt.Printf("main received from goroutine: %v\n", fromGoroutine)

	// Place value of inMain in Channel
	ch2 <- inMain
	// fromGoroutine := <-ch2
	fmt.Printf("main: %v, %v\n\n", inMain, fromGoroutine)

}
