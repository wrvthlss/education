package main

import "fmt"

type List[T any] struct {
	elements []T
}

func (l *List[T]) Add(element T) {
	l.elements = append(l.elements, element)
}

func (l *List[T]) Get(index int) T {
	return l.elements[index]
}

func main() {

	// lst instantiates a List struct. The definition
	// defines this list as a strings -- meaning all of
	// its methods align with the string type.
	lst := List[string]{}
	lst.Add("One")   // Accepts string.
	lst.Add("Two")   // Accepts string.
	lst.Add("Three") // Accepts string.

	fmt.Println(lst.Get(1)) // Accepts int, returns string.

}
