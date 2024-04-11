package main

import "fmt"

func main() {
	/* 1. */
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i, f)

	/* 2. */
	const value = 7
	var j int = value
	var k float64 = value

	fmt.Println(j, k)

	/* 3. */
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	fmt.Println(b, smallI, bigI)
}
