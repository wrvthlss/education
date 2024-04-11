package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("textFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Defer closing until return.

	// data := make([]byte, 100)
	buffer := bytes.Buffer{}
	data := make([]byte, 1024) // Temporary buffer to hold file data.

	for {
		count, err := file.Read(data)

		// Only print if count is greater than zero.
		if count > 0 {
			buffer.Write(data[:count]) // Accumulate data.
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break // Exit loop on EOF.
			}
			log.Fatal(err) // Handle non-EOF errors.
		}
	}

	// Show contents of the buffer.
	// Buffer contains the entire contents of the file.
	fmt.Println(buffer)
	fileContents := buffer.String() // Convert bytes to strings.
	fmt.Println(fileContents)       // Print contents.

	//trimData := bytes.Trim(data, "\x00")
	//byteToString := string(trimData)
	//fmt.Println(byteToString)

}
