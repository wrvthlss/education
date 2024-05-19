package main

import (
	"fmt"
	"github.com/wrvthlss/education/learning-go/ch11/pbuff/data"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/wrvthlss/education/learning-go/ch11/pbuff --go_opt=Mperson.proto=github.com/wrvthlss/education/learning-go/ch11/pbuff/data person.proto
func main() {
	p := &.Person{
		Name:  "Bob Bobson",
		Id:    20,
		Email: "bob@bobson.com",
	}
	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 data.Person
	err := proto.Unmarshal(protoBytes, &p2)
	if err != nil {
		return
	}
	fmt.Println(&p2)
}
