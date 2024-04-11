package main

import "fmt"

// Stk is generic type.
// type Stk[T any] struct {
type Stk[T comparable] struct {
	// vals is a generic member of Stk.
	// It is a slice that can contain any type of values.
	vals []T
}

// Psh is a method with a type parameter of a Stack pointer.
// This means that it can operate on Stk types.
// It takes any value -- T and pushes that value on to
// the type parameter s.
func (s *Stk[T]) Psh(val T) {
	s.vals = append(s.vals, val)
}

// PTop is a method with a type parameter of a Stack pointer.
// This means it can operate on Stk types.
// It takes no parameters, and returns the top value of the stack -- a T Value
// It returns the popped value and a boolean. The popped value is removed from the stack.
func (s *Stk[T]) PTop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	// Get the last element of the slice (the top).
	topVal := s.vals[len(s.vals)-1]
	// s.vals now contains all put the popped element (the top element is removed).
	s.vals = s.vals[:len(s.vals)-1]

	// Return the top most element, and true.
	return topVal, true
}
func (s Stk[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {

	stk := Stk[string]{}
	stk.Psh("Lorem")
	fmt.Println(stk.vals)

	fmt.Println(stk.Contains("Lorem"))

	v, ok := stk.PTop()
	fmt.Println(v, ok)

}
