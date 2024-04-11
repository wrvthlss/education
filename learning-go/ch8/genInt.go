package main

import "fmt"

// ModifyInt is an interface that has a single method defined on it.
// AddTwo returns an int.
type ModifyInt interface {
	AddTwo() int
}

// ModifyStr is an interface that has a single method defined on it.
// ConCat returns a string.
type ModifyStr interface {
	ConCat() string
}

// IntProcessor is a user defined type with underlying type of `int`.
type IntProcessor int

// AddTwo is an implementation of a function defined on ModifyInt.
// It takes the receiver `i` and add 2 to it, returning this value.
// This implicitly implements the ModifyInt interface on IntProcessor.
func (i IntProcessor) AddTwo() int {
	return int(i) + 2
}

// StrProcessor is a user defined type with underlying type of 'string'.
type StrProcessor string

// ConCat is an implementation of a function defined on ModifyStr.
// It takes the receiver `s` and concatenates a `-` and the word `cat`.
// This implicitly implements the ModifyStr interface on StrProcessor.
func (s StrProcessor) ConCat() string {
	return string(s) + "-cat"
}

/**--- Generic Function Declarations ---**/

// StringTheStrings has a type parameter of the interface ModifyStr,
// in our situation, it can only operate on StrProcessor as it is
// the only type implementing this interface.
func StringTheStrings[T ModifyStr](val T) string {
	return val.ConCat()
}

// AddToInt has a type parameter of the interface ModifyInt,
// in our situation, it can only operate on IntProcessor as it is
// the only type implementing this interface.
func AddToInt[T ModifyInt](val T) int {
	return val.AddTwo()
}

func main() {

	// Pass 3 as the value of IntProcessor
	ip := IntProcessor(3)
	// Call the AddToInt function passing in ip as the receiver to which 2 will be added.
	ipr := AddToInt(ip)
	fmt.Println("AddTwo: ", ipr)

	// Pass "Elephant" as the value of StrProcessor
	sp := StrProcessor("Elephant")
	// Call the StringTheStrings function passing in sp as the receiver to which "-cat" will be concatenated.
	spr := StringTheStrings(sp)
	fmt.Println("ConCat: ", spr)

}
