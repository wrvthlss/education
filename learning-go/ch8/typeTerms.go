package main

import (
	"errors"
	"fmt"
)

// Integer is an interface defining various type elements.
// These type elements are used to define the type which
// are assignable to type parameters and which operations
// are supported.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// divAndRemainder is a function with a type parameter T defined as an
// Integer type -- this is the interface defined above which states
// the type elements. The Integer type supports all these various
// int types. We then define the parameters -- num, denom which are of
// the generic type T which we defined as an Integer. The function then
// expects of a return of two T types and an error Type.
func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("[-] Cannot divide by zero.")
	}

	// T, T, error.
	return num / denom, num % denom, nil
}

func main() {

	// Test to see if we can successfully compile
	// uint types using the generic divAndRemainder function
	// with type elements.
	var a uint = 18_466_744_073_709_551
	var b uint = 9_223_372_036_854_775

	fmt.Println(divAndRemainder(a, b))

	// Test to see if can successfully compile int types
	// using the generic divAndRemainder function
	// with type elements.
	var myA int = 10
	var myB int = 20
	fmt.Println(divAndRemainder(myB, myA))
}
