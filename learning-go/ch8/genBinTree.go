package main

import (
	"cmp"
	"fmt"
)

// OrderFunc is a generic function signature.
// Its type parameter defines that it works on any type.
// It is a function that takes two parameters both of type T
// it returns a int. Since it returns an int this means that
// the parameters must also be ints.
type OrderFunc[T any] func(t1, t2 T) int

// Tree is a user defined, generic type.
// It has two members f, which is of type OrderFunc.
// root, which of a pointer to a generic Node type.
type Tree[T any] struct {
	f    OrderFunc[T]
	root *Node[T]
}

// Node is a generic user defined type.
// It has two members val, a generic value and
// left, right which are pointers to its containing
// struct.
type Node[T any] struct {
	val         T
	left, right *Node[T]
}

// NewTree acts as a constructor function.
// It has a generic type parameter so it operates on any type.
// As a parameter it takes a function of type OrderFunc and
// returns a pointer to a Tree type.
func NewTree[T any](f OrderFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderFunc[T], v T) bool {
	if n == nil {
		return false
	}

	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

func main() {

	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)

	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

}
