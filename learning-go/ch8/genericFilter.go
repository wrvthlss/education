package main

import "fmt"

// Filter defines a type parameter of T, accepting any type.
// It takes as parameters -- `s` a slice of type T,
// `f` a function that takes a T value and returns a Boolean,
// the Filter function itself returns a slice of type T.
func Filter[T any](s []T, f func(T) bool) []T {

	// `r` is a slice of T elements.
	var r []T

	// Get the value of `v`
	for _, v := range s {
		// If the return of f(v) is true -- the value exists.
		if f(v) {
			r = append(r, v) // If the value exist, append to r.
		}
	}

	return r
}

func main() {

	str := []string{"Lorem", "Ipsum", "Dolor", "Nurit"}

	// != "Dolor" returns false (the element does exist), meaning it
	// will not be included in the filtered slice.
	filterResult := Filter(str, func(s string) bool { return s != "Dolor" })
	fmt.Println(filterResult)

}
