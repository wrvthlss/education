package main

import "fmt"

type Innr struct {
	A int
}

func (i Innr) IntPrinter(val int) string {
	return fmt.Sprintf("Innr: %d", val)
}

func (i Innr) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outr struct {
	Innr
	S string
}

func (e Outr) IntPrinter(val int) string {
	return fmt.Sprintf("Outr: %d", val)
}

func (e Outr) StrPrinter() string {
	return fmt.Sprintf("Outr S: %s", e.S)
}

func main() {

	o := Outr{
		Innr: Innr{
			A: 10,
		},
		S: "Hello",
	}

	fmt.Println(o.Double())
	fmt.Println(o.IntPrinter(40))
	fmt.Println(o.StrPrinter())

}
