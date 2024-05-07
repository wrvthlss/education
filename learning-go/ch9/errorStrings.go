package main

import (
	"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("[!] only even numbers are processed")
	}

	return i * 2, nil
}

func doubleOdd(i int) (int, error) {
	if i%2 == 0 {
		return 0, fmt.Errorf("[-] %d is not an odd number", i)
	}

	return i * 2, nil
}

func main() {

	result, err := doubleEven(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d\n", result)
	}

	res, e := doubleOdd(2)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%d\n", res)
	}

}
