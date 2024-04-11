package main

import (
	"fmt"
	"maps"
)

func main() {

	// map[keyType]valueType
	// map literal
	keyVal := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sara", "Pete", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	fmt.Println(keyVal)

	// Using make to create a map with a default size
	// ages := make(map[int][]string, 10)

	// Write to and read from a map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Lions"])
	totalWins["Lions"]++
	totalWins["Orcas"]++
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Lions"])

	// comma ok idiom
	// To find out if a key is in a map
	m := map[string]int{
		"hello": 5,
		"world": 1,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true

	y, ok := m["world"]
	fmt.Println(y, ok) // 1 true

	z, ok := m["goodbye"]
	fmt.Println(z, ok) // 0 false

	// map comparison
	n := map[string]int{
		"hello": 5,
		"world": 1,
	}

	fmt.Println(maps.Equal(m, n)) // true
}
