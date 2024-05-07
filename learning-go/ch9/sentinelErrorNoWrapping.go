package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error.
var ErrNotFound = errors.New("not found")

func findItem(id string) error {
	return ErrNotFound
}

func main() {

	err := findItem("123")

	// Directly compare `err` to `ErrNotFound` using `==`.
	if err == ErrNotFound {
		fmt.Println("item not found")
	} else {
		fmt.Println("Item found!")
	}

}
