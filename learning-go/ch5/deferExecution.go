package main

import "fmt"

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first: ", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("second: ", val)
	}(a)

	a = 30
	defer func(val int) {
		fmt.Println("third: ", val)
	}(a)

	return a
}

func main() {

	deferExample()

}
