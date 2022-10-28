package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"hello-go-rest/internal/model/foo"

	"github.com/go-chi/chi/v5"

	log "github.com/sirupsen/logrus"
)

// Used for parsing POST/PUT JSON (no ID field)
type FooRequest struct {
	Name string `json:"name"`
}

type FooHandler struct {
	fooRepository foo.FooRepository
}

func NewFooHandler(fooRepository foo.FooRepository) *FooHandler {
	handler := new(FooHandler)
	handler.fooRepository = fooRepository
	return handler
}

func (handler *FooHandler) GetAllFoo(writer http.ResponseWriter, request *http.Request) {
	log.Debug("GetAllFoo")
	allFoo, err := handler.fooRepository.FindAll()

	if err != nil {
		handler.handleError(writer, err, 500)
		return
	}

	serializeToJson(writer, allFoo)
}

func (handler *FooHandler) GetFooById(writer http.ResponseWriter, request *http.Request) {

	fooId, err := handler.extractId(writer, request)
	if err != nil {
		return
	}

	log.Debugf("Load Foo with ID %d", fooId)

	foo, ok := handler.findFooById(writer, fooId)
	if !ok {
		return
	}

	serializeToJson(writer, foo)
}

func (handler *FooHandler) PostFoo(writer http.ResponseWriter, request *http.Request) {
	fooRequest, err := handler.parseFooRequest(writer, request)

	if err != nil {
		return
	}

	newFoo := foo.NewFoo(fooRequest.Name)
	handler.fooRepository.Save(newFoo)

	serializeToJson(writer, newFoo)
}

func (handler *FooHandler) PutFoo(writer http.ResponseWriter, request *http.Request) {
	fooRequest, err := handler.parseFooRequest(writer, request)

	if err != nil {
		return
	}

	fooId, err := handler.extractId(writer, request)
	if err != nil {
		return
	}

	existingFoo, ok := handler.findFooById(writer, fooId)
	if !ok {
		return
	}

	existingFoo.Name = fooRequest.Name

	handler.fooRepository.Save(existingFoo)

	serializeToJson(writer, &existingFoo)
}

func (handler *FooHandler) extractId(writer http.ResponseWriter, request *http.Request) (fooId int, err error) {
	id := chi.URLParam(request, "fooId")
	log.Infof("Extract ID %s", id)

	fooId, err = strconv.Atoi(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Unable to parse id %s", id)
		log.WithFields(log.Fields{"err": err}).Error(errorMessage)
		http.Error(writer, errorMessage, 400)
	}
	return fooId, err
}

func (handler *FooHandler) parseFooRequest(writer http.ResponseWriter, request *http.Request) (*FooRequest, error) {

	var fooRequest FooRequest

	err := json.NewDecoder(request.Body).Decode(&fooRequest)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Unable to parse JSON")
		http.Error(writer, err.Error(), 400)
		return nil, err
	}

	return &fooRequest, nil
}

func (handler *FooHandler) findFooById(writer http.ResponseWriter, fooId int) (*foo.Foo, bool) {
	foo, err := handler.fooRepository.FindById(fooId)

	if err != nil {
		handler.handleError(writer, err, 500)
		return nil, false
	}

	if foo == nil {
		errorMessage := fmt.Sprintf("Foo %d not found", fooId)
		log.Error(errorMessage)
		http.Error(writer, errorMessage, 404)
		return nil, false
	}

	return foo, true
}

func (handler *FooHandler) handleError(writer http.ResponseWriter, err error, code int) {
	log.Error(err.Error())
	http.Error(writer, err.Error(), 404)
}

func serializeToJson(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}
