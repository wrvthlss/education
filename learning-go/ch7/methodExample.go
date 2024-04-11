package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

/*
(p Person): Method receiver -- This defines the method on the Person type, allowing instances of Person to call this method.
String(): Method name -- By naming this method String and following the signature to return a string, it implements the `Stringer` interface from the fmt package.
string: Return type -- Specifies that this method returns a string, fulfilling the `Stringer` interface requirement.
*/
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {

	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Implicitly calls p.String() due to the fmt.Println function.
	// This is because p is an instance of Person, which implements the Stringer interface.
	fmt.Println(p)

}
