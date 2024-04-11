package main

import "fmt"

func notMut(x int) {
	x = 20
}

func mut(x *int) {
	*x = 20
}

func main() {

	x := 10

	notMut(x)
	fmt.Println(x)
	mut(&x)
	fmt.Println(x)

}
