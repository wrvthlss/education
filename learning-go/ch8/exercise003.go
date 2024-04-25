package main

import "fmt"

// Node is a struct accepting all comparable types.
// Value contains any comparable type.
// Next contains a slice of pointers to Node meaning
// Next is a pointer to the next Node of the same type,
// allowing sequential traversal.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// List is a struct accepting all comparable types.
// Head points to the first node of the list.
// Tail points to the last node of the list.
type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Add is a function with a receiver of a pointer to generic List.
// It accepts any comparable type.
func (l *List[T]) Add(t T) {
	// Insert the value passed into a Node.
	n := &Node[T]{
		Value: t,
	}

	// Add inserts a new element at the end of the list.
	// If the list is empty, it initializes both Head and
	// Tail to the new node.
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	// If Head already has a Node then insert n
	// into the last position.
	l.Tail.Next = n      // The Next node contains n
	l.Tail = l.Tail.Next // The tail node now contains n.
}

func (l *List[T]) Insert(t T, pos int) {
	// Same as line 25.
	n := &Node[T]{
		Value: t,
	}

	// Same as line 33.
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	// Check to ensure that the position parameter is valid.
	if pos <= 0 {
		// If pos is 0 or less, insert the new node at the beginning of the list.
		n.Next = l.Head
		l.Head = n
		return
	}

	// Set the current Node to the Head position.
	curNode := l.Head

	// Loop to insert values at the passed position.
	for i := 1; i < pos; i++ {
		// Find an empty node and then insert the value and return.
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		// Set the current node to the next node in the list.
		curNode = curNode.Next
	}

	// Set the currently selected node and set it as the tail.
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}

	return
}

// Index returns the index of the passed value.
func (l *List[T]) Index(t T) int {
	i := 0
	// Set the current node to the Head, loop until you find an empty node.
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		// If the value is found return the index of the value.
		if curNode.Value == t {
			return i
		}
		// Incrementer.
		i++
	}
	// If the value is not found return -1.
	return -1
}

func main() {

	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

}
