package main

import "fmt"

/*
structs cannot be modified by function parameters.
*/
type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Lorem"
}

/*
maps and slices are implemented with pointers, making
their behavior unique from structs.
*/
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	// delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
}

func main() {

	// Test call by value.
	p := person{name: "Ipsum"}
	i := 2
	s := "Hello"

	modifyFails(i, s, p)
	fmt.Println(i, s, p.name) // Not modified.

	// Mod map.
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m) // Modified.

	// Mod slice.
	q := []int{1, 2, 3}
	modSlice(q)
	fmt.Println(q) // Modified.

}
