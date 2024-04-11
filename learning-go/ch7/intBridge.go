package main

import (
	"fmt"
	"sort"
)

// Define an interface for sorting with context.
type ContextSorter interface {
	Sort([]int)
}

// Implement the interface with a type that holds state.
type DescendingSorter struct{}

func (ds DescendingSorter) Sort(slice []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
}

// A function that uses the ContextSorter interface.
func srtAndPrint(slice []int, sorter ContextSorter) {
	sorter.Sort(slice)
	fmt.Println(slice)
}

func main() {

	slice := []int{3, 1, 4, 1}
	srtAndPrint(slice, DescendingSorter{})

}
