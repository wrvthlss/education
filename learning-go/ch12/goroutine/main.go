package main

import "fmt"

func process(val int) int {
	return val * 2
}

const numGoroutines = 5

func processConcurrently(inVals []int) []int {
	// create channels
	in := make(chan int, numGoroutines)
	out := make(chan int, numGoroutines)

	// launch numGoroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for val := range in {
				result := process(val)
				out <- result
			}
		}()
	}
	// load the data into the channel in another go routine
	go func() {
		for _, v := range inVals {
			in <- v
		}
		close(in)
	}()
	// read the data
	outVals := make([]int, 0, len(inVals))
	for i := 0; i < len(inVals); i++ {
		outVals = append(outVals, <-out)

	}
	return outVals
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := processConcurrently(x)
	fmt.Println(result)
}
