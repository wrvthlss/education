package main

import (
	"errors"
	"fmt"
	"os"
)

func calRemainderAndMod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("[-] denominator is 0")
	}

	return num / den, num % den, nil
}

func main() {

	d, n, e := calRemainderAndMod(7, 0)
	fmt.Printf("%d %d %v\n", d, n, e)

	numerator := 20
	denominator := 0

	remainder, mod, err := calRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remainder, mod)

}
