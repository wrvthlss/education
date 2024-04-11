package main

import "fmt"

type Reader interface {
	Read() string
}

// FileReader is a struct implementing Reader.
type FileReader struct{}

// The Read method implicitly causes FileReader to implement Reader.
func (fr FileReader) Read() string {
	return "[+] Data from file..."
}

// NewFileReader returns a concrete FileReader struct.
func NewFileReader() FileReader {
	return FileReader{}
}

// ProcessData accepts any type that implements the Reader interface.
func ProcessData(r Reader) {
	fmt.Println(r.Read())
}

func main() {

	fr := NewFileReader()
	ProcessData(fr)

}
