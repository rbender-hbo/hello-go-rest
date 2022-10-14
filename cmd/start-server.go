package main

import (
	"log"
	"net/http"

	"hello-go-rest/internal/handler"
	"hello-go-rest/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	app := server.BuildApplication()

	fooHandler := handler.NewFooHandler(app)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/", handler.HelloWorldHandler)

	router.Route("/foo", func(router chi.Router) {
		router.Get("/", fooHandler.GetAllFooHandler)
		router.Post("/", fooHandler.PostFoo)

		router.Route("/{fooId}", func(router chi.Router) {
			router.Get("/", fooHandler.GetFooByIdHandler)
		})
	})

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
