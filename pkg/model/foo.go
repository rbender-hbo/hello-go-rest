package model

type Foo struct {
	FooId int
	Name  string
}

func NewFoo(fooId int, name string) *Foo {
	foo := new(Foo)
	foo.FooId = fooId
	foo.Name = name
	return foo
}
