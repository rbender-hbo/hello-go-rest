package main

import (
	"net/http"

	"hello-go-rest/internal/handler"
	"hello-go-rest/internal/handler2"
	"hello-go-rest/internal/rest"
	"hello-go-rest/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	log "github.com/sirupsen/logrus"
)

func main() {

	app := server.BuildApplication()
	fooHandler := handler.NewFooHandler(app.FooRepository)
	fooRestHandler := handler2.NewRestFooHandler(app.FooRepository)

	router := buildRouter(fooHandler, fooRestHandler)

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func buildRouter(fooHandler *handler.FooHandler, fooRestHandler *handler2.FooRestHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/", handler.HelloWorldHandler)
	router.Get("/hello", rest.NewHandler(handler.HelloWorld))
	router.Get("/fail", rest.NewHandler(handler.HelloFail))

	router.Route("/foo", func(router chi.Router) {
		router.Get("/", fooHandler.GetAllFoo)
		router.Post("/", fooHandler.PostFoo)

		router.Route("/{fooId}", func(router chi.Router) {
			router.Get("/", fooHandler.GetFooById)
			router.Put("/", fooHandler.PutFoo)
		})
	})

	router.Route("/foo2", func(router chi.Router) {
		router.Get("/", rest.NewHandler(fooRestHandler.GetAllFoo))
		router.Post("/", rest.NewHandler(fooRestHandler.PostFoo))

		router.Route("/{fooId}", func(router chi.Router) {
			router.Get("/", rest.NewHandler(fooRestHandler.GetFooById))
			router.Put("/", rest.NewHandler(fooRestHandler.PutFoo))
		})
	})

	return router
}
