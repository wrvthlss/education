package main

import (
	"errors"
	"fmt"
	"reflect"
)

type EmptyFieldError struct {
	FieldName string
	Err       error
}

// Implement error interface for EmptyFieldError.
func (e *EmptyFieldError) Error() string {
	return fmt.Sprintf("%s: %v", e.FieldName, e.Err)
}

var ErrEmptyField = errors.New("field cannot be empty")

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

func GetStructFields(e Employee) (reflect.Value, error) {

	// Reflect the type and value of the struct.
	pType := reflect.TypeOf(e)
	pValue := reflect.ValueOf(e)

	// Ensure provided type is a struct.
	if pType.Kind() == reflect.Struct {

		// Loop through the structs fields.
		for i := 0; i < pValue.NumField(); i++ {
			field := pType.Field(i)
			value := pValue.Field(i)

			if value.IsZero() {
				// Return custom error containing the field name.
				return value, &EmptyFieldError{
					FieldName: field.Name,
					Err:       ErrEmptyField,
				}
			}
		}
	}

	return pValue, nil

}

func main() {

	employee := Employee{
		ID:        1,
		FirstName: "",
		LastName:  "Ipsum",
	}

	_, err := GetStructFields(employee)
	var emptyFieldError *EmptyFieldError

	if errors.As(err, &emptyFieldError) {
		fmt.Println("Error: ", emptyFieldError)
	} else {
		fmt.Println("no empty fields")
	}
}
