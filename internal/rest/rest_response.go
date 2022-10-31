package rest

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Base interface for all handlers to return
type Response interface {
	Write(http.ResponseWriter)
}

// Makes code a little easier to read
// Also gives us some flexibility to change later
type HeaderMap map[string]string

// Successful response that serializes a struct to the response
type BodyResponse struct {
	Code int
	Body interface{}
	Headers HeaderMap
}


func NewResponse(body interface{}) *BodyResponse {
	response := new(BodyResponse)
	response.Body = body
	response.Code = 200
	response.Headers = make(HeaderMap)
	return response
}

func (response *BodyResponse) WithCode(code int) *BodyResponse {
	response.Code = code
	return response
}

func (response *BodyResponse) Write(writer http.ResponseWriter) {
	writer.WriteHeader(response.Code)
	writeHeaders(writer, response.Headers)
	writeBody(writer, response.Body)
}

func writeHeaders(writer http.ResponseWriter, headers HeaderMap) {
	log.Info("writeHeaders")
	for name, value := range headers {
		log.Debug("Set header ", name, " to ", value)
		writer.Header().Set(name, value)
	}
}

func writeBody(writer http.ResponseWriter, body interface{}) {
	// In the future, this could also check the request headers to determine what
	// format to serialize the body to (JSON, Protobuf, XML, etc)
	log.Info("Write Body", body)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(body)
}
