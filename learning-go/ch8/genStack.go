package main

import "fmt"

// Stak is a user defined, generic type which
// functions on any underlying type. It has a single
// member, items, which a slice data structure
// that can contain any type.
type Stak[T any] struct {
	items []T
}

// Push is a generic method defined on the
// type Stak, meaning it functions on Stak types.
// It accepts a generic parameter, item.
func (s *Stak[T]) Push(item T) {
	s.items = append(s.items, item)
}

func main() {

	s := Stak[int]{}
	s.Push(10)
	fmt.Println(s.items)

}
