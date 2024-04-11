package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindEmployee(name string) bool {

	if name == m.Name {
		fmt.Println("[+] Manager.")
	} else {
		for _, report := range m.Reports {
			if report.Name == name {
				fmt.Println("[+] Employee:", name, "found!")
				return true
			}
		}
	}

	fmt.Println("[-] Employee not found...")
	return false
}

func main() {

	m := Manager{
		Employee: Employee{
			Name: "Lorem Ipsum",
			ID:   "12345",
		},
		Reports: []Employee{
			{Name: "Alice", ID: "2984"},
			{Name: "Bob", ID: "5783"},
		},
	}

	fmt.Println(m.FindEmployee("Alice"))

}
