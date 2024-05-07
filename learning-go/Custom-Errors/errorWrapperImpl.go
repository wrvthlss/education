package main

import (
	"errors"
	"fmt"
)

// Access will generate the second error if flag is
// set to 1.
func Access(usr, pwd string, flag int) (string, error) {
	if flag == 1 {
		err := errors.New("error flag set to on")
		return "", fmt.Errorf("error in access(2): %w", err)
	}

	// Enter generates the third error.
	usr, err := Enter(usr, pwd)
	if err != nil {
		return "", fmt.Errorf("error in access(3): %w", err)
	}

	return "welcome to the system", nil
}

func Enter(usr, pwd string) (string, error) {
	if usr == "admin" && pwd == "password" {
		return "user:admin", nil
	}
	// Third error message set and sent back to Access.
	err := errors.New("invalid credentials")

	// Fourth error message set if third is set.
	return "", fmt.Errorf("error in enter(4): %w", err)
}

func main() {

	res, err := Access("admin", "admin", 0)
	// Generate first error.
	if err != nil {
		fmt.Println("Error(1):", err)
		return
	}

	fmt.Println(res)
}
