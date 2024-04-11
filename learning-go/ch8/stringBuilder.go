package main

import "fmt"

type StringBuilder string

func (sb StringBuilder) Append(str string) StringBuilder {
	return StringBuilder(string(sb) + str)
}

func (sb StringBuilder) Display() {
	fmt.Println(sb)
}

type Addition int

func (add Addition) Plus(y int) Addition {
	return Addition(int(add) + y)
}

func (add Addition) DoubleResult() Addition {
	i := int(add) * 2
	return Addition(i)
}

func (add Addition) Result() {
	fmt.Println(add)
}

func main() {

	sb := StringBuilder("Hello")
	sb.Append(" World").Display()

	add := Addition(3)
	add.Plus(5).DoubleResult().Result()
}
