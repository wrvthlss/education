package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	example "github.com/wrvthlss/education/learning-go/ch11/protobufExample/generated"
	"log"
)

func main() {
	// Create a new person
	person := &example.Person{
		Name:  "John Doe",
		Id:    12345,
		Email: "johndoe@example.com",
	}

	// Serialize the person to a binary format
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Marshaling error: ", err)
	}

	// Print the binary data
	fmt.Println("Serialized data: ", data)

	// Deserialize the binary data back to a person object
	newPerson := &example.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	// Print the deserialized person
	fmt.Printf("Deserialized Person: %+v\n", newPerson)
}
