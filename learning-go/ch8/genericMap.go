package main

import "fmt"

// Map has two type parameter -- T1 and T2 which can be any type.
// It takes two parameters -- s, a slice of T1 values,
// a function which takes a T1 value and returns a T2 value.
// The Map function, itself, returns a slice of T2 values.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	// Create a slice, r, with a capacity and length of argument s.
	r := make([]T2, len(s))

	// `i` is the index, and `v` is the value of the `s` slice.
	//  r[i] -- the index of r is equal to...
	// f(v) -- the value returned by 'f'.
	// This effectively places all values returned by 'f' into the slice 'r'
	for i, v := range s {
		r[i] = f(v)
	}

	// return 'r'
	return r
}

func passFunc(str string, f func(s string) int) int {
	lgnth := f(str)

	return lgnth
}

func square(x int) int { return x * x }

func getLen(str string) int { return len(str) }

func main() {

	num := []int{3, 5, 7, 9}

	m := Map(num, square)
	fmt.Println(m)

	w := passFunc("Beginning", getLen)
	fmt.Println(w)

}
