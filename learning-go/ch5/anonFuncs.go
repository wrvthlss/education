package main

import "fmt"

/*
Declare package scope variables assigned anonymous functions.
*/
var (
	addition = func(i, j int) int { return i + j }
	subtract = func(i, j int) int { return i - j }
)

/*
Define new functions within a function and assign it to A variable.
*/
func main() {

	// Define anonymous function.
	// Declare an anonymous function with the keyword `func` immediately
	// followed by the input parameters, the return values, and the opening brace.
	f := func(j int) {
		fmt.Println("printing", j, "from inside an anonymous function")
	}

	// Iterate over i and pass its value to anonymous function `f()`
	for i := 0; i < 5; i++ {
		f(i)
	}

	// Use package variables.
	x := addition(2, 3)
	fmt.Println(x)

}
