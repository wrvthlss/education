package main

import "fmt"

func f1(a string) int {
	return len(a)
}

// Sum the unicode character value.
func f2(a string) int {
	total := 0

	// _ == index
	// v == value
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main() {

	// Define function signature for the variable
	// to be the same as original function.
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result)

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)

}
