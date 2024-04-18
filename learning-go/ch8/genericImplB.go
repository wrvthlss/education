package main

import (
	"errors"
	"fmt"
)

type Value int

type Calculator interface {
	Sum(Value, Value) int
	Prod(Value, Value) int
	DvR(Value, Value) (int, int, error)
}

// SimpleCalc is a new type with.
type SimpleCalc struct {
	sum  func(Value, Value) int
	prod func(Value, Value) int
	dvr  func(Value, Value) (int, int, error)
}

func (s *SimpleCalc) Sum(x, y Value) int {
	return s.sum(x, y)
}

func (s *SimpleCalc) Prod(x, y Value) int {
	return s.prod(x, y)
}

// DvR delegates the function to SimpleCalc dvr.
// dvr already matches the expected return types specified
// in the interface. The return types are implied to match
// exactly between the interface definition and the method
// implementation.
func (s *SimpleCalc) DvR(x, y Value) (int, int, error) {
	return s.dvr(x, y)
}

func sum(a, b Value) int {
	return int(a + b)
}

func prod(a, b Value) int {
	return int(a * b)
}

func dvr(a, b Value) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("[-] division by zero not allowed")
	}
	return int(a / b), int(a % b), nil
}

func main() {

	calc := SimpleCalc{
		sum:  sum,
		prod: prod,
		dvr:  dvr,
	}

	fmt.Println(calc.Sum(10, 5))
}
