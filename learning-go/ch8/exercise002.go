package main

import (
	"fmt"
	"strconv"
)

// Printable is a generic interface that implements fmt.Stringer
//
// Incorrect:
// Printable has an underlying type of int or float64, meaning
// the receivers must be of type int or float.
//
// Printable is defined with a type constraint that limits its
// implementers to those who underlying types are `int` or
// `float64`. Printable is an interface, not a type with an
// underlying type.
type Printable interface {
	fmt.Stringer
	~int | ~float64 // Only types with underlying types of int or float64 can implement Printable.
}

// PrintInt and FloatInt are user defined types with underlying types,
// int and float64 which are the required underlying types for
// implementations of Printable.
type PrintInt int

type PrintFloat float64

// PrntInt and PrntFlt implement the String method, this causes
// the two methods to implicitly implement the Printable
// interface.
func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%.2f", pf)
}

// PrintIt takes any Printable and prints its string representations.
// Since PrintInt and PrintFloat implement fmt.Stringer due to the
// String method, they conform the Printable interface.
func PrintIt[T Printable](t T) {
	fmt.Println(t) // Automatically calls the String method.
}

func main() {

	var i PrintInt = 3
	PrintIt(i)

	var y PrintFloat = 5.0
	PrintIt(y)

}
