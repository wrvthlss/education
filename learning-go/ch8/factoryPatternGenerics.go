package main

import "fmt"

// Animal is an interface for animal types.
type Animal interface {
	Speak() string
}

// Dog is a type that satisfies the Animal interface.
type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

// Cat is another type that satisfies the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

// AnimalFactory is a generic function that returns a new instance of an Animal.
func AnimalFactory[T Animal]() T {
	var animal T
	return animal
}

func main() {

	dog := AnimalFactory[Dog]()
	cat := AnimalFactory[Cat]()

	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())

}
