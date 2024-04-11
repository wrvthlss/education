package main

import "fmt"

func main() {

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {

		// `size` is scoped to the entire switch statement.
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, " -- short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, " -- medium word -- ", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, " -- long word --")
		}
	}

	// Breaking out of a nested switch.
loop:
	for i := 0; i < 10; i++ {

		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "[+] Even")
		case 3:
			fmt.Println(i, "[+] Divisible by 3, but not 2.")
		case 7:
			fmt.Println("[!] Exiting the loop...", "\n")
			break loop
		default:
			fmt.Println(i, "[-] Boring...")
		}
	}

	// Blank switch statement.
	//     A regular `switch` statement only allows you to check a
	//     value for equality. A blank `switch` allows you to use
	// 	   any boolean comparison for each case.
	tokens := []string{"hi", "salutations", "hello"}
	for _, word := range tokens {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "[!] short word...")
		case wordLen > 10:
			fmt.Println(word, "[!] long word...")
		default:
			fmt.Println(word, "[+] correct length!")
		}
	}

}
