package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(idKey)
	fmt.Println(nameKey)

	//x = x + 1 // Error: cannot assign to x
	//y = "bye" // Error: cannot assign to y

	fmt.Println(x)
	fmt.Println(y)

}
