package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(x, y int) (int, error) {
	return x + y, nil
}

func sub(x, y int) (int, error) {
	return x - y, nil
}

func mul(x, y int) (int, error) {
	return x * y, nil
}

func dvd(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("[-] Division by zero.")
	}

	return x / y, nil
}

var opMaps = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": dvd,
}

func main() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"6", "/", "3"},
		{"6", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("[-] Invalid expression: ", expression)
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]

		opFunc, ok := opMaps[op]
		if !ok {
			fmt.Println("[-] Unsupported operator: ", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)

	}
}
