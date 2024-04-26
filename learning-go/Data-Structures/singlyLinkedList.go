package main

import "fmt"

// Node is a generic type accepting any value with a Value field containing
// values of type `T`, and Next, a pointer to the next Node in the list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// List is a generic type accepting any value with a Head, which is a pointer
// to the Head node.
type List[T any] struct {
	Head *Node[T]
}

// AddToFront takes a val and populates a new node's Value field with it.
// The method then sets the next node to the current Head. It then
// sets the newNode to be the Head.
func (l *List[T]) AddToFront(val T) {
	newNode := &Node[T]{
		Value: val,
		Next:  l.Head, // This is the current head node.
	}

	l.Head = newNode // The newly inserted node is now the new head node.
}

// Print starts at the current Head and traverses current.Next until
// nil is reached, signalling the end of the list.
func (l *List[T]) Print() {
	for current := l.Head; current != nil; current = current.Next {
		fmt.Printf("%v -> ", current.Value)
	}

	fmt.Println("nil")
}

func main() {

	list := &List[int]{}
	list.AddToFront(3)
	list.AddToFront(5)
	list.AddToFront(7)
	list.AddToFront(9)

	list.Print()
}
