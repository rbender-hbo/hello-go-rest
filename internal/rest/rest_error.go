package rest

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code int
	Message string
	Headers HeaderMap
}

type errorBody struct {
	Message string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return NewErrorResponseWithMessage(err.Error())
}

func NewErrorResponseWithMessage(message string, a...any) *ErrorResponse {
	formattedMessage := fmt.Sprintf(message, a...)
	response := new(ErrorResponse)
	response.Code = 500
	response.Message = formattedMessage
	response.Headers = make(HeaderMap)
	return response
}

func (response *ErrorResponse) Write(writer http.ResponseWriter) {

	log.Error(response.Message)

	writer.WriteHeader(response.Code)
	writeHeaders(writer, response.Headers)

	body := errorBody{Message: response.Message}
	writeBody(writer, body)
}

func (response *ErrorResponse) WithCode(code int) *ErrorResponse {
	response.Code = code
	return response
}
