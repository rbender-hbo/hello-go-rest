package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"hello-go-rest/internal/model/foo"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestFooHandlerGetAllFoo(t *testing.T) {

	req, _ := http.NewRequest("GET", "/foo", nil)

	recorder := httptest.NewRecorder()

	fooRepository := buildFooRepository()
	foo1, _ := fooRepository.FindById(1)
	foo2, _ := fooRepository.FindById(2)

	fooHandler := NewFooHandler(fooRepository)
	handler := http.HandlerFunc(fooHandler.GetAllFoo)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)

	var parsedFoos []foo.Foo
	err := json.NewDecoder(recorder.Body).Decode(&parsedFoos)
	assert.Nil(t, err)

	assert.Contains(t, parsedFoos, *foo1)
	assert.Contains(t, parsedFoos, *foo2)
}

func TestFooHandlerGetFooById(t *testing.T) {

	req, _ := http.NewRequest("GET", "/foo/1", nil)
	req = addUrlParam(req, "fooId", "1")

	recorder := httptest.NewRecorder()

	fooRepository := buildFooRepository()
	foo1, _ := fooRepository.FindById(1)

	fooHandler := NewFooHandler(fooRepository)
	handler := http.HandlerFunc(fooHandler.GetFooById)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)

	parsedFoo := parseFoo(t, recorder)
	assert.Equal(t, parsedFoo, *foo1)
}

func TestFooHandlerGetFooByIdReturns404WhenNotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/foo/3", nil)
	req = addUrlParam(req, "fooId", "3")

	recorder := httptest.NewRecorder()

	fooHandler := buildFooHandler()
	handler := http.HandlerFunc(fooHandler.GetFooById)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

func TestFooHandlerPostFoo(t *testing.T) {

	requestBody := strings.NewReader(`{"name":"PostFoo"}`)
	req, _ := http.NewRequest("POST", "/foo", requestBody)

	recorder := httptest.NewRecorder()

	fooRepository := buildFooRepository()
	fooHandler := NewFooHandler(fooRepository)
	handler := http.HandlerFunc(fooHandler.PostFoo)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)

	parsedFoo := parseFoo(t, recorder)
	assert.Equal(t, fooRepository.LatestFoo.FooId, parsedFoo.FooId)
	assert.Equal(t, "PostFoo", parsedFoo.Name)
}

func TestFooHandlerPutFoo(t *testing.T) {

	requestBody := strings.NewReader(`{"name":"PutFoo"}`)
	req, _ := http.NewRequest("PUT", "/foo/1", requestBody)
	req = addUrlParam(req, "fooId", "1")

	recorder := httptest.NewRecorder()

	fooHandler := buildFooHandler()
	handler := http.HandlerFunc(fooHandler.PutFoo)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)

	parsedFoo := parseFoo(t, recorder)
	assert.Equal(t, 1, parsedFoo.FooId)
	assert.Equal(t, "PutFoo", parsedFoo.Name)
}

func TestFooHandlerPutFooReturns404WhenNotFound(t *testing.T) {

	requestBody := strings.NewReader(`{"name":"PutFoo"}`)
	req, _ := http.NewRequest("PUT", "/foo/3", requestBody)
	req = addUrlParam(req, "fooId", "3")

	recorder := httptest.NewRecorder()

	fooHandler := buildFooHandler()
	handler := http.HandlerFunc(fooHandler.PutFoo)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

func addUrlParam(request *http.Request, name string, value string) *http.Request {
	routeContext := chi.NewRouteContext()
	routeContext.URLParams.Add(name, value)
	requestContext := context.WithValue(request.Context(), chi.RouteCtxKey, routeContext)
	return request.WithContext(requestContext)
}

func parseFoo(t *testing.T, recorder *httptest.ResponseRecorder) foo.Foo {
	var parsedFoo foo.Foo
	err := json.NewDecoder(recorder.Body).Decode(&parsedFoo)
	assert.Nil(t, err)
	return parsedFoo
}

func buildFooHandler() FooHandler {
	fooRepo := buildFooRepository()
	return *NewFooHandler(fooRepo)
}

func buildFooRepository() *foo.InMemoryFooRepository {
	foo1 := foo.NewFooWithId(1, "One")
	foo2 := foo.NewFooWithId(2, "Two")

	fooRepo := foo.NewInMemoryFooRepository()
	fooRepo.Save(foo1)
	fooRepo.Save(foo2)

	return fooRepo
}
