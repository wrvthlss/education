package main

import (
	"fmt"
	"strconv"
)

// Stringer is an interface for types that can be converted to a string.
type Stringer interface {
	ToString() string
}

// Inter is an interface for types that can be converted to an int.
type Inter interface {
	ToInt() int
}

type IntWrapper int

func (i IntWrapper) ToString() string {
	return strconv.Itoa(int(i))
}

func (i IntWrapper) ToInt() int {
	return int(i)
}

// FloatWrapper wraps a float, but only implements Stringer.
type FloatWrapper float64

func (f FloatWrapper) ToString() string {
	return fmt.Sprintf("%f", f)
}

/* Generic functions to utilize the interfaces. */
func ConvertToString[T Stringer](value T) string {
	return value.ToString()
}

func ConvertToInt[T Inter](value T) int {
	return value.ToInt()
}

func main() {
	iw := IntWrapper(42)
	fmt.Println("String: ", ConvertToString(iw))
	fmt.Println("Int: ", ConvertToInt(iw))

	fw := FloatWrapper(3.14)
	fmt.Println("String: ", ConvertToString(fw))
}
