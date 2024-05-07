package main

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p Person) error {

	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}

	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}

	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil

}

func main() {

	prsn := Person{
		FirstName: "Lorem",
		LastName:  "",
		Age:       50,
	}

	res := ValidatePerson(prsn)
	fmt.Println(res)

}
