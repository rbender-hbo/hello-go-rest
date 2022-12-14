package foo

type Foo struct {
	FooId int    `json:"fooId"`
	Name  string `json:"name"`
}

func NewFoo(name string) *Foo {
	foo := new(Foo)
	foo.Name = name
	return foo
}

func NewFooWithId(fooId int, name string) *Foo {
	foo := new(Foo)
	foo.FooId = fooId
	foo.Name = name
	return foo
}
