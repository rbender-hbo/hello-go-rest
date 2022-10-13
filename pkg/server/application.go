package server

/*
Data structure to hold various objects that can be injected into each other.

TODO: Replace with real Dependency Injection framework (dig, wire, etc)
*/

import (
	"hello-go-rest/pkg/model/foo"
)

type Application struct {
	FooRepository *foo.FooRepository
}

func BuildApplication() *Application {

	app := new(Application)
	app.FooRepository = buildFooRepository()

	return app
}

func buildFooRepository() *foo.FooRepository {
	fooRepository := foo.NewFooRepository()

	fooRepository.Save(foo.NewFoo(1, "FooOne"))
	fooRepository.Save(foo.NewFoo(2, "FooTwo"))
	fooRepository.Save(foo.NewFoo(3, "FooThree"))

	return fooRepository
}
