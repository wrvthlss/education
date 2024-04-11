package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	// Return the top (last added) element of the stack.
	top := s.vals[len(s.vals)-1]
	// Remove the top element from the stack.
	s.vals = s.vals[:len(s.vals)-1]

	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {

		if v == val {
			return true
		}
	}
	return false
}

func main() {
	// Ensure you define your stack using `var` and the type `int`.
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	v, ok := intStack.Pop()
	fmt.Println(v, ok)

	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(5))
	fmt.Println(intStack.Contains(30)) // Value was popped.

}
