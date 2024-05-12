package main

import (
	"errors"
	"fmt"
	"strings"
)

// MultiError is a custom error type that holds multiple errors
type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	// Build a single string with all error message concatenated
	var sb strings.Builder
	for _, err := range m.Errors {
		sb.WriteString(err.Error() + "\n")
	}

	return sb.String()
}

// Add appends an error to the MultiError list
func (m *MultiError) Add(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

func (m *MultiError) IsEmpty() bool {
	return len(m.Errors) == 0
}

var AuthorizationError = errors.New("authorization error -- check credentials")
var AccessError = errors.New("access error -- check error flag")

type AuthError struct {
	Err  error
	Code int
}

type AccError struct {
	Err  error
	Code int
}

func (a *AuthError) Error() string {
	return fmt.Sprintf("%s: %v", a.Err, a.Code)
}

func (acc *AccError) Error() string {
	return fmt.Sprintf("%s: %v", acc.Err, acc.Code)
}

func Access(usr, pwd string, flag int) (string, error) {
	multiErr := &MultiError{}

	if flag == 1 {
		err := AccError{
			Err:  AccessError,
			Code: 1001,
		}
		multiErr.Add(fmt.Errorf("error accessing system(b) -- Access(): %w", err.Err))
	}

	usr, err := Enter(usr, pwd)
	if err != nil {
		multiErr.Add(fmt.Errorf("error in authentication(c) -- Enter(): %w", err))
	}

	if multiErr.IsEmpty() {
		return "access granted", nil
	}

	return "", multiErr
}

func Enter(usr, pwd string) (string, error) {
	if usr == "admin" && pwd == "password" {
		return "user:admin", nil
	}

	err := AuthError{
		Err:  AuthorizationError,
		Code: 1002,
	}

	return "", fmt.Errorf("error in authentication(d) -- Enter(): %w", err.Err)
}

func main() {

	res, err := Access("admin", "passwrd", 0)

	if err != nil {
		fmt.Println("Errors(a) -- main():")
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
