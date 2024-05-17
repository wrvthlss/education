package main

import (
	"fmt"
	"github.com/wrvthlss/education/learning-go/ch10/typeAlias/fubar"
)

func main() {

	// Remember, for fields within structs to be exported,
	// they must also be capitalized.
	foo := fubar.Foo{
		X: 2,
		S: "Lorem",
	}
	fmt.Println(foo.Contents())
	fmt.Printf("%v \n\n", foo)

	// type alias
	fmt.Println("using type alias--")
	type Bar = fubar.Foo
	bar := Bar{
		X: 3,
		S: "Ipsum",
	}
	fmt.Println(bar.Contents())
	fmt.Println(bar)

}
