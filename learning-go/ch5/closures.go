package main

import "fmt"

/*
Functions declared inside of functions are special; they are closures.
Functions declared inside functions are able to access and modify
variables declared in the outer function.
*/

func main() {
	a := 20
	b := 50

	f := func() {
		fmt.Println(a) // 20
		a = 30         // assign value to a '='.
	}

	j := func() {
		fmt.Println(b) // 50
		b := 30        // create a new `b` that is destroyed when closure exits ':='.
		fmt.Println(b) // 30
	}

	f()
	fmt.Println(a, "\n") // 30

	j()
	fmt.Println(b) // 50
}
