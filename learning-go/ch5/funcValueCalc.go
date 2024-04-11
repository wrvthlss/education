package main

import (
	"fmt"
	"strconv"
)

/*
	Building a simple calculator using functions as values in a map.
*/

type opFuncType func(int, int) int

func ad(i int, j int) int { return i + j }
func sb(i int, j int) int { return i - j }
func ml(i int, j int) int { return i * j }
func dv(i int, j int) int { return i / j }

var opMap = map[string]opFuncType{
	"+": ad,
	"-": sb,
	"*": ml,
	"/": dv,
}

func main() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"9", "/", "3"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("[-] Invalid expression: ", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Print(result)
	}

}
