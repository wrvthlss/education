package main

import "fmt"

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}

	return num / denom
}

func divAndRemainder(num, denom int) (int, int) {
	if denom == 0 {
		return 0, 0
	}

	return num / denom, num % denom
}

// MyFuncOpts
// Simulate named and optional function parameters using structs.
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	if opts.Age > 18 {
		fmt.Println("[+] You may enter...")
	} else {
		fmt.Println("[!] Must be older to continue.")
	}

	return nil
}

func addTo(base int, vals ...int) []int {
	// Define a slice with no elements and capacity based on the length of vals.
	out := make([]int, 0, len(vals))

	// _ -- no index
	// v -- value
	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func main() {

	result := div(6, 2)
	fmt.Println(result)

	agePass := MyFuncOpts{Age: 30}
	ageFail := MyFuncOpts{Age: 17}

	MyFunc(ageFail)
	MyFunc(agePass)

	// Your base number is 3 and your value. Additional values are
	// added to this base number and then stored as the next element.
	b := addTo(3, 4, 5, 6, 7) // [7, 8, 9, 10]
	fmt.Println(b)

	// Supply two variables to capture quotient and remainder.
	divResult, divRem := divAndRemainder(7, 2)
	fmt.Print(divResult, ", ") // quotient
	fmt.Print(divRem)          // remainder

}
