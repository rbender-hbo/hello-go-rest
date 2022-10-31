package handler2

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hello-go-rest/internal/model/foo"
	"hello-go-rest/internal/rest"

	"github.com/go-chi/chi/v5"

	log "github.com/sirupsen/logrus"
)

// Used for parsing POST/PUT JSON (no ID field)
type FooRequest struct {
	Name string `json:"name"`
}

type FooRestHandler struct {
	fooRepository foo.FooRepository
}

func NewRestFooHandler(fooRepository foo.FooRepository) *FooRestHandler {
	handler := new(FooRestHandler)
	handler.fooRepository = fooRepository
	return handler
}

func (handler *FooRestHandler) GetAllFoo(request *http.Request) rest.Response {
	log.Debug("GetAllFoo")
	allFoo, err := handler.fooRepository.FindAll()

	if err != nil {
		return rest.NewErrorResponse(err)
	}

	return rest.NewResponse(allFoo)
}

func (handler *FooRestHandler) GetFooById(request *http.Request) rest.Response {

	fooId, err := handler.extractId(request)
	if err != nil {
		return err
	}

	log.Debugf("Load Foo with ID %d", fooId)

	foo, err := handler.findFooById(fooId)
	if err != nil {
		return rest.NewErrorResponseWithMessage("Could not find foo %d", fooId).WithCode(404)
	}

	return rest.NewResponse(foo)
}

func (handler *FooRestHandler) PostFoo(request *http.Request) rest.Response {
	fooRequest, err := handler.parseFooRequest(request)

	if err != nil {
		return err
	}

	newFoo := foo.NewFoo(fooRequest.Name)
	handler.fooRepository.Save(newFoo)

	return rest.NewResponse(newFoo)
}

func (handler *FooRestHandler) PutFoo(request *http.Request) rest.Response {
	fooRequest, err := handler.parseFooRequest(request)

	if err != nil {
		return err
	}

	fooId, err := handler.extractId(request)
	if err != nil {
		return err
	}

	existingFoo, err := handler.findFooById(fooId)
	if err != nil {
		return rest.NewErrorResponseWithMessage("Could not find Foo %d", fooId).WithCode(404)
	}

	existingFoo.Name = fooRequest.Name

	handler.fooRepository.Save(existingFoo)

	return rest.NewResponse(existingFoo)
}

func (handler *FooRestHandler) extractId(request *http.Request) (int, *rest.ErrorResponse) {
	id := chi.URLParam(request, "fooId")
	log.Infof("Extract ID %s", id)

	fooId, err := strconv.Atoi(id)
	if err != nil {
		return 0, rest.NewErrorResponseWithMessage("unable to parse id %s", id)
	}

	return fooId, nil
}

func (handler *FooRestHandler) parseFooRequest(request *http.Request) (*FooRequest, *rest.ErrorResponse) {

	var fooRequest FooRequest

	err := json.NewDecoder(request.Body).Decode(&fooRequest)
	if err != nil {
		return nil, rest.NewErrorResponse(err)
	}

	return &fooRequest, nil
}

func (handler *FooRestHandler) findFooById(fooId int) (*foo.Foo, *rest.ErrorResponse) {
	foo, err := handler.fooRepository.FindById(fooId)

	if err != nil {
		return nil, rest.NewErrorResponse(err)
	}

	if (foo == nil) {
		return nil, rest.NewErrorResponseWithMessage("Could not find Foo %d", fooId).WithCode(404)
	}

	return foo, nil
}
