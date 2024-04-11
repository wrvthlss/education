package main

import "fmt"

func updateSlice(strSlice []string, str string) {
	strSlice[len(strSlice)-1] = str
	fmt.Println(strSlice)
}

// Using `append` to change the length is not reflected in the original variable.
func growSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str) // Append only affects the copy.
	fmt.Println(strSlice)            // Shows additional element.
}

func main() {

	x := []string{"Lorem", "Ipsum", "Dolor", "Nurit"}
	fmt.Println(x)

	updateSlice(x, "[+]")
	fmt.Println(x)

	growSlice(x, "[-]") // Append only affects the copy.
	fmt.Println(x)      // Additional element not reflected in original.

}
