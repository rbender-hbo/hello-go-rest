package main

import (
	"net/http"

	"hello-go-rest/internal/handler"
	"hello-go-rest/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	log "github.com/sirupsen/logrus"
)

func main() {

	app := server.BuildApplication()
	fooHandler := handler.NewFooHandler(app.FooRepository)
	productHandler := handler.NewProductHandler(app.ProductService)

	router := buildRouter(fooHandler, productHandler)

	log.Println("Starting server :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func buildRouter(fooHandler *handler.FooHandler, productHandler *handler.ProductHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/", handler.HelloWorldHandler)

	router.Route("/foo", func(router chi.Router) {
		router.Get("/", fooHandler.GetAllFoo)
		router.Post("/", fooHandler.PostFoo)

		router.Route("/{fooId}", func(router chi.Router) {
			router.Get("/", fooHandler.GetFooById)
			router.Put("/", fooHandler.PutFoo)
		})
	})

	router.Get("/product/{productId}", productHandler.GetProduct)

	return router
}
