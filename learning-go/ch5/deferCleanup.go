package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// The first value in `os.Args` is the name of the program.
	if len(os.Args) < 2 {
		log.Fatal("[!] No file specified...")
	}

	// The first argument supplied is the name of the file to read from.
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Normally, a function call runs immediately. However, `defer`
	// delays the invocation until the surrounding function exits.
	// The code within `defer` runs after the return statement.
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Variable to hold file contents.
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)    // Read file contents in bytes.
		os.Stdout.Write(data[:count]) // Write file contents to stdout.

		if err != nil {
			if err != io.EOF { // Check for end of file marker.
				log.Fatal(err) // All other errors.
			}
			break // EOF
		}
	}

}
