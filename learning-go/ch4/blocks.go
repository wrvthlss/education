package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Shadowing variables
	x := 10

	if x > 5 {
		fmt.Println(x) // 10

		// Shadow x
		x := 5

		fmt.Println(x) // 5, shadowing variable
	} // end shadowing variable scope

	fmt.Println(x)    // 10
	fmt.Println("\n") // 10

	// Shadowing multiple assignment
	y := 10

	if y > 5 {
		y, z := 5, 20 // shadow y
		fmt.Println(y, z)
	}
	fmt.Println(y)

	fmt.Println("\n")

	// if
	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("[-] Too low...: ", n)
	} else if n > 5 {
		fmt.Println("[-] Too big...: ", n)
	} else {
		fmt.Println("[+] Correct: ", n)
	}

	fmt.Println(n)

	fmt.Println("\n")

	// Scoping a variable to an if statement
	if m := rand.Intn(10); m == 0 {
		fmt.Println("[!] Too low...: ", m)
	} else if n > 5 {
		fmt.Println("[!] Too big...: ", m)
	} else {
		fmt.Println("[+] Correct!: ", m)
	}

	// fmt.Println(m) -- error scoped to if statement.

	// for statement
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("\n")
		}
	}

	// Initialization based on value calculated before loop
	i := 0
	for ; i < 3; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("\n")
		}
	}

	// Increment rules defined in the loop
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			i++
			fmt.Println(i)
		} else {
			i += 2
			fmt.Println(i)
		}
	}
	fmt.Println("\n")

	// Confusing code
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		if i%5 == 0 {
	//			fmt.Println("FizzBuzz")
	//		} else {
	//			fmt.Println("Fizz")
	//		}
	//	} else if i%5 == 0 {
	//		fmt.Println("Buzz")
	//	} else {
	//		fmt.Println(i)
	//	}
	//}

	// Confusing code less confusing using continue
	// Idiomatic Go encourages short if statement bodies.
	for i := 1; i < 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("[+] FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("[+] Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("[+] Buzz")
			continue
		}

		fmt.Println(i)
	}

}
