package server

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
