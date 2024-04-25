package main

import "fmt"

type ValidTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Doubler[T ValidTypes](a T) T {
	return a * 2
}

func main() {
	dblr := Doubler(7)
	fmt.Println(dblr)
}
