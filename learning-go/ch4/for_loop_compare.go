package main

import "fmt"

func main() {

	/*
		Both iterate over the second through the second to last elements in an array.
		Skip over the first element and print up to the penultimate element -- len(evenVals)-1
	*/
	// `for-range` loop
	evenVals := []int{2, 4, 6, 8, 10}

	for i, v := range evenVals {
		if i == 0 {
			continue
		}

		if i == len(evenVals)-1 {
			break
		}

		fmt.Println(i, v)
	}

	fmt.Println("\n")

	// standard for loop
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}

}
