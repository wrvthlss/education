package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	fmt.Println(x)

	// First two elements
	y := x[:2]
	fmt.Println(y)

	// Skip first element
	z := x[1:]
	fmt.Println(z)

	// Elements between 1 and 3, non-inclusive -- [b,c]
	d := x[1:3]
	fmt.Println(d)

	// Copy entire slice
	e := x[:]
	fmt.Println(e)

	fmt.Println("\n")

	// Full slice expression
	xx := make([]string, 0, 5)
	xx = append(xx, "a", "b", "c", "d")
	yy := xx[:2:2]
	zz := xx[2:4:4]
	fmt.Println(yy)
	fmt.Println(zz)
}
