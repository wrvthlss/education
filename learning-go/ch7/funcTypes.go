package main

import (
	"fmt"
	"sort"
)

// Function type that takes a slice of integers and returns nothing.
type Sorter func([]int)

// A function that matches the `Sorter` function type.
func ascendingOrder(slice []int) {
	sort.Ints(slice)
}

// A function using the `Sorter` function type as a parameter.
func sortAndPrint(slice []int, sorter Sorter) {
	sorter(slice)
	fmt.Println(slice)
}

func main() {
	slice := []int{3, 1, 4, 1}
	sortAndPrint(slice, ascendingOrder) // Sorts in ascending order and prints.
}
