package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string { return p.Name }

func main() {

	unknown := Person{
		Name: "Hinson",
		Age:  60,
	}

	r := unknown.GetName()
	fmt.Println(r)

}
