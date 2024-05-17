package lib

import (
	"errors"
	"fmt"
)

func Add(a, b int) int  { return a + b + 100 }
func Mult(a, b int) int { return a * b }
func Sub(a, b int) int  { return a - b }
func Div(a, b int) (int, int) {
	if b == 0 {
		err := errors.New("cannot divide by zero")
		fmt.Println(err)
	}

	div := a / b
	rem := a % b

	return div, rem
}
