package server

/*
Data structure to hold various objects that can be injected into each other.

TODO: Replace with real Dependency Injection framework (dig, wire, etc)
*/

import (
	"hello-go-rest/internal/model/foo"
	"hello-go-rest/internal/service"
	"os"
)

type Application struct {
	FooRepository foo.FooRepository
	ProductService service.ProductService
}

func BuildApplication() *Application {

	app := new(Application)
	app.FooRepository = buildFooRepository()

	baseURL := getEnvOrDefault("PRODUCT_BASE_URL", "https://dummyjson.com")

	app.ProductService = *service.NewProductService(baseURL)

	return app
}

func getEnvOrDefault(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func buildFooRepository() foo.FooRepository {
	fooRepository := foo.NewInMemoryFooRepository()

	fooRepository.Save(foo.NewFooWithId(1, "FooOne"))
	fooRepository.Save(foo.NewFooWithId(2, "FooTwo"))
	fooRepository.Save(foo.NewFooWithId(3, "FooThree"))

	return fooRepository
}
