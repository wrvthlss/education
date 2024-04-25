package main

import "fmt"

/*
Suppose you have an existing function that performs a basic calculation on a set of numbers.
Let's say this function calculates the sum of the numbers. If a new requirement comes up—for instance,
you now also need to calculate the average after summing the numbers—you can handle this new
requirement through function delegation.
*/

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func CalculateAverage(numbers []int) float64 {
	total := Sum(numbers)
	return Average(total, len(numbers)) // Delegate new functionality.
}

func Average(total, count int) float64 {
	if count == 0 {
		return 0 // Avoid division by zero.
	}

	return float64(total) / float64(count)
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Println(result)

	avg := CalculateAverage(numbers)
	fmt.Println(avg)

}
