package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var data = `
a: Hello
b:
    c: [1, 2]
`

type Info struct {
	A string
	B struct {
		C []int
	}
}

func main() {
	info := Info{}

	err := yaml.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", info.A)
}
