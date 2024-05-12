package main

import (
	"fmt"

	format "github.com/wrvthlss/education/do-format"
	"github.com/wrvthlss/education/math"
)

func main() {
	num := math.Double(2)
	output := format.Numbers(num)

	fmt.Println(output)
}
