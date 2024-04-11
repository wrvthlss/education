package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(str string) (int, error) {
	f, err := os.Open(str)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fileSize, _ := f.Stat()

	return int(fileSize.Size()), nil
}

func fileLenB(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	data := make([]byte, 2040)
	total := 0

	for {
		count, err := f.Read(data)
		total += count

		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return total, nil
}

func main() {

	i, err := fileLen("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

	count, e := fileLen("output.txt")
	if err != nil {
		log.Fatal(e)
	}

	fmt.Println(count)

}
