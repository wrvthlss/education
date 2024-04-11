package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {

	myAdder := Adder{start: 10}
	fmt.Println(myAdder.start)

	// Assigning a method to a variable  -- method value.
	f1 := myAdder.AddTo
	fmt.Println(f1(5))

	// Create a function from the type itself -- method expression.
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 20)) // First parameter is the receiver for the method.

}
