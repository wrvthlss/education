package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return prefix + " " + body
	}
}

func manipInt(x int) func(int) int {
	return func(add int) int {
		return x + add
	}
}

func main() {

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Lorem"))
	fmt.Println(helloPrefix("Ipsum"))

	num := manipInt(7)
	fmt.Println(num(2))
	fmt.Println(num(4))
	fmt.Println(num(6))
	fmt.Println(num(10))

}
