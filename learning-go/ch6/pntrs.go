package main

import "fmt"

func main() {
	var x int32 = 10
	var y bool = true

	ten := 10
	var pointerToTen *int
	pointerToTen = &ten
	fmt.Println(pointerToTen)
	fmt.Println(*pointerToTen)

	// & -- get memory address
	pointerX := &x
	pointerY := &y

	fmt.Print("[!] Address for x: ", "\t", pointerX, "\n", "[!] Address for y: ", "\t", pointerY, "\n\n")

	fmt.Println("[!] Value for pointerX: ", *pointerX)

	// Assigning primitive variables does not share memory space.
	j := 10           // Set j to 10.
	i := j            // Set i to j.
	i = 20            // Change i to see if changes are reflected in j (shared memory address?).
	fmt.Println(i, j) // The change in i is not reflected in j.
}
