package main

import "fmt"

const (
	field1 = 0
	field2 = 1 + iota // iota has val of 1
	field3 = 20
	field4        // 20, no type or value explicitly assigned -- copies line 8
	field5 = iota // 4, fifth line and iota starts at 0
)

func main() {

	fmt.Println(field1, field2, field3, field4, field5)

}
