package main

import "fmt"

func div60(i int) {
	// Register the function with defer to handle
	// a potential panic.
	defer func() {
		// Recover and check if a non-nil value was found.
		// recover must be called in defer, with panics
		// only deferred functions are ran.
		if v := recover(); v != nil { // v stores the panic message, check not nil to check for message.
			fmt.Println(v) // print panic message.
		}
	}()
	fmt.Println(60 / i)
}

func main() {

	// The 3rd element (0) will cause a panic to run.
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

}
