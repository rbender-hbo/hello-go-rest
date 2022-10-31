package rest

import "net/http"

type RestHandlerFunc func(request *http.Request) Response

func NewHandler(handler RestHandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response := handler(request)
		response.Write(writer)
	}
}
