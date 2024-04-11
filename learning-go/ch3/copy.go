package main

import "fmt"

func main() {
	// array
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)

	// y defined with length of 2
	copy(y, d[:])
	fmt.Println(y) // [5 6]

	// d defined with length of 4
	copy(d[:], x)
	fmt.Println(d) //[1 2 3 4]

	// Convert array to slice
	xArray := [4]int{5, 6, 7, 8} // array {}
	xSlice := xArray[:]          // converts entire array to slice []

	fmt.Println(xSlice)
}
