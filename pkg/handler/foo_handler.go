package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"hello-go-rest/pkg/model/foo"
	"hello-go-rest/pkg/server"

	"github.com/go-chi/chi/v5"
)

type FooHandler struct {
	fooRepository *foo.FooRepository
}

func NewFooHandler(app *server.Application) *FooHandler {
	handler := new(FooHandler)
	handler.fooRepository = app.FooRepository
	return handler
}

func (handler *FooHandler) GetAllFooHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	allFoo := handler.fooRepository.FindAll()
	serializeToJson(writer, allFoo)
}

func (handler *FooHandler) GetFooByIdHandler(writer http.ResponseWriter, request *http.Request) {

	fooId, err := handler.extractId(request)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	log.Printf("Load Foo with ID %d", fooId)

	foo, ok := handler.fooRepository.FindById(fooId)
	if !ok {
		errorMessage := fmt.Sprintf("Foo %d not found", fooId)
		http.Error(writer, errorMessage, 404)
		return
	}

	serializeToJson(writer, foo)
}

func (handler *FooHandler) extractId(request *http.Request) (fooId int, err error) {
	//id := request.URL.Path[len("/foo/"):]
	id := chi.URLParam(request, "fooId")
	log.Printf("Extract ID %s", id)

	fooId, err = strconv.Atoi(id)
	return fooId, err
}

func serializeToJson(writer http.ResponseWriter, data interface{}) {
	json.NewEncoder(writer).Encode(data)
}
