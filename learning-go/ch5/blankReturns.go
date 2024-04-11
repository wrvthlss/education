package main

import (
	"errors"
	"fmt"
)

// Named returns with blank return
func divideAndRemain(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("[-] Cannot divide by zero.")
		return
	}

	result, remainder = num/denom, num%denom
	return
}

func main() {

	quo, rem, _ := divideAndRemain(6, 4)
	fmt.Println(quo, rem)

}
