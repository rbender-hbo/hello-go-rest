package main

import (
	"log"
	"net/http"

	"hello-go-rest/pkg/handler"
	"hello-go-rest/pkg/model/foo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	fooRepository := foo.NewFooRepository()
	fooRepository.Save(foo.NewFoo(1, "FooOne"))
	fooRepository.Save(foo.NewFoo(2, "FooTwo"))
	fooRepository.Save(foo.NewFoo(3, "FooThree"))

	fooHandler := handler.NewFooHandler(fooRepository)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/", handler.HelloWorldHandler)

	router.Route("/foo", func(router chi.Router) {
		router.Get("/", fooHandler.AllFooHandler)

		router.Route("/{fooId}", func(router chi.Router) {
			router.Get("/", fooHandler.FooByIdHandler)
		})
	})

	//http.HandleFunc("/", handler.HelloWorldHandler)
	//http.HandleFunc("/foo", fooHandler.AllFooHandler)
	//http.HandleFunc("/foo/", fooHandler.FooByIdHandler)

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
