package main

import (
	"fmt"
	"slices"
)

/*
	Using [...] makes an Array.
	Using [] manes a Slice.
*/

func main() {

	// Arrays
	var x [3]int
	var y = [...]int{10, 20, 30}

	fmt.Println(x[0])
	fmt.Println(y[0])

	// Slices
	// slice literal
	var j = []int{10, 20, 30} // do not define the size of a slice

	i := []int{1, 2, 3, 4, 5}
	k := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}

	fmt.Println(j[1])
	fmt.Println(slices.Equal(i, k)) // true
	fmt.Println(slices.Equal(i, z)) // false
	// fmt.Println(slices.Equal(i, s)) // will not compile

	w := append(i, 10)
	z = append(z, 7)
	fmt.Println(w[0])
	fmt.Println(w[5])
	fmt.Println(z[6])

	fmt.Println("\n")

	// Using make to specify type, length, and capacity (opt.)
	xx := make([]int, 5)
	fmt.Println(xx, len(xx), cap(xx)) // [0,0,0,0,0] 5 5

	// Defining capacity with make
	yy := make([]int, 5, 10)
	fmt.Println(yy, len(yy), cap(yy)) // [0,0,0,0,0] 5 10

	// Declaring a slice with default values
	data := []int{2, 4, 6, 8}
	fmt.Println(data)

	// Using make with 0 length and defined capacity
	// This allows the use of append, if the number of
	// items is smaller, no extra 0s will be added.
	// If the number is larger, the code will not panic.
	zz := make([]int, 0, 10)
	zz = append(zz, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println(zz)

}
