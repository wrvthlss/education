package main

import "fmt"

// Operations is a generic interface with two methods.
type Operations interface {
	ShowString() string
	ShowInt() int
}

// Vals is a new type containing a field storing a string
// and a field storing an int.
type Vals struct {
	strValue string
	intValue int
}

// ShowString defined on Vals.
func (sV Vals) ShowString() string {
	return sV.strValue
}

// ShowInt defined on Vals.
func (iV Vals) ShowInt() int {
	return iV.intValue
}

/*
Now that ShowString and ShowInt are defined on Vals, Vals now has implicitly
implemented the Operations interface.
*/

func main() {

	intStore := Vals{intValue: 4}
	strStore := Vals{strValue: "Lorem"}

	fmt.Println(strStore.ShowString())
	fmt.Println(intStore.ShowInt())

}
