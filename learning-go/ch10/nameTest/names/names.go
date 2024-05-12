package names

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

var EmptyStringError = errors.New("string error: string is empty")

type EmpStrError struct {
	Err  error
	Code string // String for error code to provide more context.
}

func (es *EmpStrError) Error() string {
	return fmt.Sprintf("%s: %v", es.Err, es.Code)
}

var StringLengthError = errors.New("string error: extraction requires at least 5 characters")

type StrLenError struct {
	Err  error
	Code string // String for error code to provide more context.
}

func (sle *StrLenError) Error() string {
	return fmt.Sprintf("%s: %v", sle.Err, sle.Code)
}

func Extract(s string) (string, error) {
	if utf8.RuneCountInString(s) == 0 {
		return "", fmt.Errorf("%w", &EmpStrError{
			Err:  EmptyStringError,
			Code: "errorcode: 1",
		})
	}

	if utf8.RuneCountInString(s) < 5 {
		return "", fmt.Errorf("%w", &StrLenError{
			Err:  StringLengthError,
			Code: "errorcode: 2",
		})
	}

	return string([]rune(s)[:5]), nil
}

// Format converts a string to it's uppercase version.
// Initial error handling is done in Extract, if Extract
// generates no errors, this function is safe.
func Format(s string) string { return strings.ToUpper(s) }
