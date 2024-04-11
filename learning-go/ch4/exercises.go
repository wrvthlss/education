package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// 1. Place 100 random numbers into an `int` slice.

	// randNums := []int{}
	randNums := make([]int, 100)
	for i := 0; i < 100; i++ {
		randNums = append(randNums, rand.Intn(100))
	}
	fmt.Println(randNums)

	// 2.  Apply rules based on values of a random slice.
	for i := 1; i < len(randNums); i++ {

		if i%2 == 0 {
			fmt.Println(i, " is divisible by two!")
		}

		if i%3 == 0 {
			fmt.Println(i, " is divisible by three!")
		}

		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, " is divisible by 6!")
		}

		fmt.Println("[-] Nevermind...")

	}
	fmt.Println("\n")

	// 2.  Apply rules based on values of a random slice -- using `for/range` and `switch`.
	for _, v := range randNums {

		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Nevermind")
		}

	}

	// Variable scope.
	var total int
	for i := 0; i < 10; i++ {

		if i == 0 {
			fmt.Println("[!] In for loop...")
			continue
		}

		// total := total + i -- Places the new value in `total.
		total = total + i // Applies summation and changes outer `total`.
		fmt.Println(total)
	}
	fmt.Println("\n", "[!] Out of for loop: ", total)

}
