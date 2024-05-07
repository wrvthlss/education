package main

import (
	"errors"
	"fmt"
)

var InvalidID = errors.New("invalid ID")

func GetUserID(i int) (int, error) {
	if i == 1337 {
		return 1, nil
	}

	return 0, InvalidID
}

func main() {

	_, err := GetUserID(2112)
	if err != nil {
		if errors.Is(err, InvalidID) {
			fmt.Println("the provided ID is invalid")
		}
	}

}
