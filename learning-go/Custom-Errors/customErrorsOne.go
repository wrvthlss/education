package main

import (
	"fmt"
)

// StatusCode is user-defined type with an underlying type of int.
type StatusCode int

// Define error codes.
const (
	InvalidCredentials = StatusCode(1)
	ConnectionFailed   = StatusCode(2)
)

// NetworkError is a struct containing Code and Message.
// Code is defined as a StatusCode type and Message, a string.
type NetworkError struct {
	Code    StatusCode
	Message string
}

// NetworkError implements the Error method, meaning it implicitly
// implements the error interface.
func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error [%d]: %s", e.Code, e.Message)
}

// Connect takes two strings and an int, and returns string and error.
// The flag is set to trigger the NetworkError or not.
func Connect(usr, pwd string, flag int) (string, error) {
	// If the flag is set to 0, trigger the NetworkError.
	if flag == 0 {
		return "", &NetworkError{
			Code:    ConnectionFailed,
			Message: "connection failed",
		}
	}

	// Delegate to Login function.
	usr, err := Login(usr, pwd)
	if err != nil {
		return "", err
	}

	return "welcome to the system", nil
}

// Login takes two strings and returns a string and an error.
func Login(usr, pwd string) (string, error) {
	if usr == "admin" && pwd == "password" {
		return "usr:admin", nil
	}

	// If the credentials fail, trigger NetworkError and
	// set the Code to InvalidCredentials.
	return "", &NetworkError{
		Code:    InvalidCredentials,
		Message: "invalid credentials",
	}

}

func main() {

	usr, err := Connect("admin", "admin", 1)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(usr)

}
