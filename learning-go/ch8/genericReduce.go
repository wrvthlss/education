package main

import "fmt"

// Reduce has T1 and T2 defined as type parameters that can be of any type.
// It takes -- `s` a slice of T1 values, `init` a T2 value
// `f` a function that takes a T1 and T2 parameter and returns a T2.
// The Reduce function, itself, returns a T2.
func Reduce[T1, T2 any](s []T1, init T2, f func(T2, T1) T2) T2 {

	r := init

	for _, v := range s {
		// r = 0 + 3 = 3
		// r = 3 + 5 = 8
		// r = 8 + 7 = 15
		// r = 15 + 9 = 24
		r = f(r, v)
	}

	return r
}

func addY(x, y int) int { return x + y }

func main() {

	vals := []int{3, 5, 7, 9} // Line 14.

	r := Reduce(vals, 0, addY)

	fmt.Println(r)

}
