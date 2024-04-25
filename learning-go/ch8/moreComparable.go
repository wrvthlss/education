package main

import "fmt"

type Comp[T comparable] struct {
	vals []T
}

type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("[+] ThingInt: ", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("[+] ThingSlice: ")
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("[+] equal")
	}
}

// CompareSlices is a generic function comparing two slices of any comparable type.
func CompareSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func main() {

	// test
	tInt := ThingerInt(2)
	tIntA := ThingerInt(2)
	// fmt.Println(tInt)
	Comparer(tInt, tIntA) // Comparable, matching types.

	// sInt := []int{1, 3, 5, 7, 9}
	// fmt.Println(sInt)

	// var a int = 10
	// var b int = 10
	// Comparer(a, b) // comparable

	// Cannot compare when underlying types match.
	// Comparer(a, tInt)

	// You can also not compare ThingerSlice or []int
	// sInt_a := []int{1,3,5,7,9}
	// Comparer(sInt, sInt_a)

	// It is legal to call Comparer on types of `Thinger` our interface.
	// You can also compare it to method implementing the interface.
	var c Thinger = tInt
	var d Thinger = tIntA
	Comparer(c, d)

	stringsA := []string{"one", "two", "three"}
	stringsB := []string{"one", "two", "three"}
	fmt.Println(CompareSlices(stringsA, stringsB))

	intsA := []int{1, 3, 5, 7, 9}
	intsB := []int{1, 3, 5, 7, 9}
	fmt.Println(CompareSlices(intsB, intsA))

}
