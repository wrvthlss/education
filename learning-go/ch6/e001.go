package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makePerson(firstname, lastname string, age int) Person {

	return Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}

}

func makePersonPointer(firstname, lastname string, age int) *Person {

	return &Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}

}

func main() {

	fn := "Lorem"
	ln := "Ipsum"
	a := 50

	makePersonPointer(fn, ln, a)
	makePerson(fn, ln, a)

}
