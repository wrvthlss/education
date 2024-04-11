package main

import "fmt"

func main() {

	// Print sub-slices.
	greeting := []string{"Hello", "Hola", "Velcome", "Welcome"}
	sl1 := greeting[0:2]
	sl2 := greeting[1:3]
	sl3 := greeting[2:4]
	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)

	// Print runes.
	message := "Hi 🤦‍♀️and 🤦‍♂️"
	runes := []rune(message)      // String to runes.
	fmt.Println(runes)            // Unicode values.
	fmt.Println(string(runes[3])) // Cast 4th rune to string.

	// Define struct and instantiate.
	type Employee struct {
		firstname string
		lastname  string
		id        int
	}

	Lorem := Employee{
		id: 1098343,
	}
	fmt.Println(Lorem.id)

	Dolor := Employee{
		"Dolor",
		"Nurit",
		9883239,
	}
	fmt.Println(Dolor)

	var emp = Employee{}
	emp.firstname = "Pricipium"
	emp.lastname = "Rosundit"
	emp.id = 769302
	fmt.Println(emp)
}
