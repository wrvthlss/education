package main

import (
	"errors"
	"fmt"
)

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
	if flag == 1 {
		err := AccError{
			Err:  AccessError,
			Code: 1001,
		}
		return "", fmt.Errorf("error accessing system(b) -- Access(): %w", err.Err)
	}

	usr, err := Enter(usr, pwd)
	if err != nil {
		return "", fmt.Errorf("error in authentication(c) -- Enter(): %w", err)
	}

	return "access granted", nil
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
		fmt.Println("error(a) -- main(): ", err)
	}

	fmt.Println(res)

}
