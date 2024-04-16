package main

type OrderFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

// NewTree acts as a constructor function.
func NewTree[T any](f OrderFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

// TODO: Complete tree.

func main() {

	// TODO: Implement.

}
