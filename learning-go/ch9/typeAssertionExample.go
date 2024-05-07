package main

import (
	"fmt"
	"os"
)

func mightFail() error {
	return &os.PathError{
		Op:   "open",
		Path: "/invalid/path/to/file",
		Err:  os.ErrNotExist,
	}
}

func main() {

	err := mightFail()
	if serr, ok := err.(*os.PathError); ok {
		fmt.Printf("Failed to %s %s: %v\n", serr.Op, serr.Path, serr.Err)
	} else {
		fmt.Println("Generic error: ", err)
	}

}
