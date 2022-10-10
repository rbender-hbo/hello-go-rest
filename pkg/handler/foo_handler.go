package handler

import (
	"encoding/json"
	"net/http"

	"hello-go-rest/pkg/model"
)

type FooHandler struct {
	fooRepository *model.FooRepository
}

func NewFooHandler(fooRepository *model.FooRepository) *FooHandler {
	handler := new(FooHandler)
	handler.fooRepository = fooRepository
	return handler
}

func (handler *FooHandler) AllFooHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	allFoo := handler.fooRepository.FindAll()
	json.NewEncoder(writer).Encode(allFoo)
}
