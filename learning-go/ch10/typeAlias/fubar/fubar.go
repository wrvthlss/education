package fubar

type Foo struct {
	X int
	S string
}

func (f Foo) Hello() string   { return "Hello" }
func (f Foo) Goodbye() string { return "Good-bye" }
func (f Foo) Contents() Foo {
	return Foo{
		X: 1,
		S: "String",
	}
}
