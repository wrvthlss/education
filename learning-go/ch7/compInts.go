package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {

	di := DoubleInt(3)
	di.Double()
	fmt.Println(di)

	dis := DoubleIntSlice{3, 5, 7}
	dis.Double()
	fmt.Println(dis)

	var tdi DoubleInt = 10
	var tdi2 DoubleInt = 10
	DoublerCompare(&tdi, &tdi2)

	result := tdi.Double
	result()
	fmt.Println(tdi)

}
