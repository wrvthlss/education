package main

import "fmt"

type Inner struct {
	X int
}

type Outer struct {
	Inner
	Y int
}

type Calculation struct {
	Outer
	Modifier int
}

func (c Calculation) ModifyCalculation() int {
	return c.X + c.Y + c.Modifier
}

func main() {

	j := Outer{
		Inner{X: 3},
		5,
	}

	c := Calculation{
		Outer:    j,
		Modifier: 7,
	}

	r := c.ModifyCalculation()
	fmt.Println(r)

}
