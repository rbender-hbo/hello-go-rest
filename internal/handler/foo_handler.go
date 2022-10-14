package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"hello-go-rest/internal/model/foo"
	"hello-go-rest/internal/server"

	"github.com/go-chi/chi/v5"

	log "github.com/sirupsen/logrus"
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
	log.Debug("GetAllFoo")
	allFoo := handler.fooRepository.FindAll()
	serializeToJson(writer, allFoo)
}

func (handler *FooHandler) GetFooByIdHandler(writer http.ResponseWriter, request *http.Request) {

	fooId, err := handler.extractId(request)
	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	log.Debug("Load Foo with ID %d", fooId)

	foo, ok := handler.fooRepository.FindById(fooId)
	if !ok {
		errorMessage := fmt.Sprintf("Foo %d not found", fooId)
		log.Error(errorMessage)
		http.Error(writer, errorMessage, 404)
		return
	}

	serializeToJson(writer, foo)
}

func (handler *FooHandler) PostFoo(writer http.ResponseWriter, request *http.Request) {
	var newFoo foo.Foo

	err := json.NewDecoder(request.Body).Decode(&newFoo)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("Unable to parse JSON")
		http.Error(writer, err.Error(), 400)
		return
	}

	handler.fooRepository.Save(&newFoo)

	serializeToJson(writer, &newFoo)
}

func (handler *FooHandler) extractId(request *http.Request) (fooId int, err error) {
	id := chi.URLParam(request, "fooId")
	log.Printf("Extract ID %s", id)

	fooId, err = strconv.Atoi(id)
	return fooId, err
}

func serializeToJson(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}
