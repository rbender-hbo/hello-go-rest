package handler

import (
	"errors"
	"hello-go-rest/internal/rest"
	"net/http"
)

type Hello struct {
	Greeting string `json:"name"`
}

func HelloWorld(request *http.Request) rest.Response {
	hello := new(Hello)
	hello.Greeting = "Hey Everybody!"

	return rest.NewResponse(hello)
}

func HelloFail(request *http.Request) rest.Response {

	err := errors.New("FAIL")
	return rest.NewErrorResponse(err)
}
