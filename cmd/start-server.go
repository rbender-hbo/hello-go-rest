package main

import (
	"log"
	"net/http"

	"hello-go-rest/pkg/handler"
	"hello-go-rest/pkg/model/foo"
)

func main() {

	fooRepository := foo.NewFooRepository()
	fooRepository.Save(foo.NewFoo(1, "FooOne"))
	fooRepository.Save(foo.NewFoo(2, "FooTwo"))
	fooRepository.Save(foo.NewFoo(3, "FooThree"))

	fooHandler := handler.NewFooHandler(fooRepository)

	http.HandleFunc("/", handler.HelloWorldHandler)
	http.HandleFunc("/foo", fooHandler.AllFooHandler)
	http.HandleFunc("/foo/", fooHandler.FooByIdHandler)

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
