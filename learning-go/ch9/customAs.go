package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) CodeVals() []int {
	return me.Codes
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v ", me.Codes)
}

func AFunctionThatReturnsAnError() error {
	return MyErr{Codes: []int{1, 3, 5, 7, 9, 11}}
}

func main() {

	// Returns our slice of error codes.
	err := AFunctionThatReturnsAnError()

	// Use var to declare a variable of a specific type set
	// to the zero value.
	var myErr MyErr

	// If err, is of the same type as the pointer to myErr,
	// return the codes.
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Codes)
	}

	// coder defines CodeVals() as a member, which if of
	// type MyErr.
	var coder interface {
		CodeVals() []int
	}

	// This works because coder contains CodeVals(),
	// a type of MyErr. This will give us the same return
	// as errors.As(err, &myErr).
	if errors.As(err, &coder) {
		fmt.Println(coder.CodeVals())
	}

}
