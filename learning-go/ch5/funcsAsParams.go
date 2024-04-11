package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func sendMessage(mess string) func(string) string {
	// Return function that concatenates `reply` to `mess`.
	return func(reply string) string {
		return mess + reply
	}
}

func main() {

	people := []Person{
		{"Pat", "Patterson", 47},
		{"Tracy", "Borem", 33},
		{"Fred", "Fredson", 28},
	}

	// Sort the slice by LastName.
	// `people` is "captured by" the closure.
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)

	// Sort the slice by Age.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

	// Returning functions from functions.
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	m := sendMessage("Hello")
	fmt.Println(m(" [+] Reply"))

}
