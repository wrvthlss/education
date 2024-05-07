package main

import (
	"errors"
	"fmt"
)

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

// Is compares instances of ResourceErr.
func (re ResourceErr) Is(target error) bool {
	// Check if target is of type ResourceErr via type assertion.
	if other, ok := target.(ResourceErr); ok {
		// If target is a ResourceErr, compare various combination of the
		// fields to determine if the current error `re` matches the
		// target.
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code

		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}

	return false
}

func main() {

	err := ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	// Match err to a ResourceErr. This matches because the
	// resource is provided and the code field is ignored.
	if errors.Is(err, ResourceErr{Resource: "Database"}) {
		fmt.Println("database error: ", err)
	}

	// Same as above, except using err2.
	if errors.Is(err2, ResourceErr{Resource: "Network"}) {
		fmt.Println("network error: ", err2)
	}

	// Matches err to a ResourceErr. This matches because the
	// code is provided and resource field is ignored.
	if errors.Is(err, ResourceErr{Code: 123}) {
		fmt.Println("code 123 error: ", err)
	}

	// Matches err fully against ResourceErr, both Resource and
	// Code are provided.
	if errors.Is(err, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered:", err)
	}

	// Does not match, because err has different values than the passed
	// resource. We are passing the Database resource error and comparing
	// it against the Network resource errors.
	if errors.Is(err, ResourceErr{Resource: "Network", Code: 456}) {
		fmt.Println("Network is broken and Code 123 triggered:", err)
	}
}
