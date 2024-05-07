package main

import (
	"errors"
	"fmt"
	"slices"
)

// MyErr is a custom error type.
type MyErr struct {
	Codes []int
}

// Error causes MyErr to implement the error interface.
func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

// Is is a custom Is implementation. It allows to define
// how to compare two errors of type `MyErr`.
func (me MyErr) Is(target error) bool {
	// Check whether target matches MyErr by comparing
	// their Codes slices. Assert value of MyErr.
	if me2, ok := target.(MyErr); ok {
		// Compares the slices and return true if both slices match.
		return slices.Equal(me.Codes, me2.Codes)
	}

	return false
}

func compareErrors(err1, err2 error) {
	if errors.Is(err1, err2) {
		fmt.Println("the errors are equivalent")
	} else {
		fmt.Println("the errors are different")
	}
}

func main() {

	// Two custom errors with the same slice values.
	err1 := MyErr{Codes: []int{1, 2, 3}}
	err2 := MyErr{Codes: []int{1, 2, 3}}

	// Error with different values.
	err3 := MyErr{Codes: []int{4, 5, 6}}

	// compareErrors checks for error equivalence.
	fmt.Println("compare err1 and err2: ")
	compareErrors(err1, err2)

	// compareErrors for non-matching
	compareErrors(err1, err3)

}
