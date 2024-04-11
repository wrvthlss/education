package main

import "fmt"

func main() {
	var x int = 10
	x *= 2

	fmt.Println(x)

	// Returns Unicode code points.
	var myFirstInitial rune = 'J'
	var myLastInitial int32 = 'B'
	fmt.Println(myFirstInitial, myLastInitial)

	// Returns strings.
	var myFirstName string = "John"
	var myLastName string = "Doe"
	fmt.Println(myFirstName, myLastName)

	// Explicit type conversion.
	var z int = 10
	var y float64 = 30.2

	// Convert z to float64 and add it to y.
	var sum1 float64 = float64(z) + y

	// Convert y to int and add it to z.
	var sum2 int = z + int(y)

	fmt.Println(sum1, sum2)

	// Variables with different types.
	var i, j = 10, "hello"
	fmt.Println(i, j)
}
