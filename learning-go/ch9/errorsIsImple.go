package main

import (
	"errors"
	"fmt"
	"os"
)

func fChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// Error is wrapped.
		return fmt.Errorf("in fileChecker: %w", err)
	}

	f.Close()
	return nil
}

func main() {

	err := fChecker("not_here.txt")
	if err != nil {
		// errors.Is checks for the error based on type.
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file does not exist.")
		}
	}

}
