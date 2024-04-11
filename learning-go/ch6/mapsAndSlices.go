package main

import "fmt"

func addTwo(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] + 2 // Adding 2 to the original slice elements.
	}
}

func mapManip(m map[string]int) {
	for k, v := range m {
		m[k] = v + 2 // Adding 2 to the original map values.
	}
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	addTwo(s)      // Pass slice to function.
	fmt.Println(s) // Test for mutability.

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	mapManip(m)    // Pass map to function.
	fmt.Println(m) // Test for mutability.

}
