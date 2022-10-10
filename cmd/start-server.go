package main

import (
	"log"
	"net/http"

	"hello-go-rest/pkg/handler"
	"hello-go-rest/pkg/model"
)

func main() {

	fooRepository := model.NewFooRepository()
	fooRepository.Save(model.NewFoo(1, "FooOne"))
	fooRepository.Save(model.NewFoo(2, "FooTwo"))
	fooRepository.Save(model.NewFoo(3, "FooThree"))

	fooHandler := handler.NewFooHandler(fooRepository)

	http.HandleFunc("/", handler.HelloWorldHandler)
	http.HandleFunc("/foo", fooHandler.AllFooHandler)

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
