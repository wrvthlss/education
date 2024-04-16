package main

import "fmt"

// OrderableFunc is a generic type alias for a function type.
// It is important to note here, that the return type T,
// must be the same underlying type as the parameter(s) T.
type OrderableFunc[T any] func(t1, t2 T) T

func intCompare(t1, t2 int) int {
	return t1 - t2
}

func strCompare(s1, s2 string) string {
	if s1 == s2 {
		return "true"
	}

	return "false"
}

func main() {

	// Declaring a variable of type OrderableFunc[int] and [string].
	var compare OrderableFunc[int]
	var sCompare OrderableFunc[string]

	// Assign intCompare to compare, since it matches the function signature.
	compare = intCompare
	sCompare = strCompare

	// Use compare function.
	fmt.Println(compare(5, 10))
	fmt.Println(sCompare("abc", "abc"))

}
