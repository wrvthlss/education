package main

import (
	"cmp"
	"fmt"
)

// Min returns the minimum of two ordered values.
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func main() {

	fmt.Println(Min(8, 3))
	fmt.Println(Min("apple", "banana"))
	fmt.Println(Min(3.14, 1.59))

}
