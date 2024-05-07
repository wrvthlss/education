package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
}

func (l *List[T]) AddToFront(val T) {
	newNode := &Node[T]{
		Value: val,
		Next:  l.Head,
	}

	l.Head = newNode
}

func (l *List[T]) Traverse() {
	for current := l.Head; current != nil; current = current.Next {
		fmt.Printf("%v --> ", current.Value)
	}

	fmt.Println("nil")
}

func main() {

	list := List[int]{}
	list.AddToFront(3)
	list.AddToFront(5)
	list.AddToFront(7)
	list.AddToFront(9)

	list.Traverse()

}
