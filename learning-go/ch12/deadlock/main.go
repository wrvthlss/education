package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine // Write 1 into ch1 -- send
		fromMain := <-ch2  // Read value from ch2 and assign to fromMain -- receive
		fmt.Println("goroutine: ", inGoroutine, fromMain)
	}()

	inMain := 2
	var fromGoroutine int

	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}

	fmt.Println("main: ", inMain, fromGoroutine)

}
