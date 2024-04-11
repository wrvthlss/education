package main

import (
	"errors"
	"fmt"
)

func divAndRemainders(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("[-] Cannot divide by zero")
		return num, denom, err
	}
	result, remainder = num/denom, num%denom

	return result, remainder, err
}

func main() {

	quotient, remainder, _ := divAndRemainders(9, 4)
	fmt.Println("Quotient: ", quotient, "\t", "Remainder: ", remainder)

	fmt.Println()

	quotient, remainder, err := divAndRemainders(8, 0)
	fmt.Println(err, "\t", "Numerator: ", quotient, "\t", "Denominator: ", remainder)
}
