package main

import "fmt"

func main() {

	// slice
	evenVals := []int{2, 4, 6, 8, 10, 12}

	// i -- index
	// v -- value
	for i, v := range evenVals {
		fmt.Println(i, v) // (0,2),(1,4),. . .,(5,12)
	}

	fmt.Println("\n")

	// slice range no position
	for _, v := range evenVals {
		fmt.Println(v)
	}

	fmt.Println("\n")

	// Key and no value
	// Using `map` to define a set()
	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// `map` iteration.
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	fmt.Println("\n")

	// Embedded loop -- keys and values have random orders.
	for i := 0; i < 3; i++ {
		// Loop 0,1,2.
		fmt.Println("Loop", i)

		for k, v := range m {
			fmt.Println(k, v)
		}

	}

	fmt.Println("\n")

	// String iteration
	samples := []string{"hello", "apple"}
	for _, sample := range samples { // Extract the string.

		for i, r := range sample { // Iterate over the string.
			fmt.Println(i, r, string(r)) // r -- unicode, string(r) -- char
		}

		fmt.Println() // "\n"
	}

	fmt.Println()

	outerValues := []int{1, 3, 5, 7, 9, 13, 15, 17, 19}
	innerValues := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}

	// Experimenting with labels.
outer:
	for _, outerVal := range outerValues {

		fmt.Println("[!] outer val: ", outerVal)

		for _, innerVal := range innerValues {
			// process inner val
			fmt.Println(innerVal)

			if innerVal%3 == 0 {
				fmt.Println("[!] Modulo 3.")
				continue outer
			}
		}
		fmt.Println("[+] Completed loop...")

	}
}
