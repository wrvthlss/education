package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	bob := person{
		"Bob",
		40,
		"dog",
	}

	// struct literal
	// Fields must be assigned in order
	// All fields must be assigned a value
	julia := person{
		"Julia",
		40,
		"cat",
	}

	// struct literal
	// Fields can be defined in any order
	// Values are not required for all fields
	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Println(julia.name)
	fmt.Println(beth.age)
	fmt.Println(bob.pet)
}
