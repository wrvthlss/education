package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func FindEmployee(employees []Employee, id int) (string, bool) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp.Name, true // Found.
		}
	}
	return "[-] Employee not found.", false // Not found, return error message and false.
}

func main() {

	employees := []Employee{
		{
			ID:   1,
			Name: "Lorem Ipsum",
		},
		{
			ID:   2,
			Name: "Dolor Nurit",
		},
	}

	// Return values (found): emp, true
	// Return values (not found): emp, false
	emp, ok := FindEmployee(employees, 1)

	if ok { // ok == True
		fmt.Println("[+] Found employee: ", emp)
	} else { // ok == False
		fmt.Println("[-] Employee not found...")
	}

}
