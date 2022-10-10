package handler

import (
	"fmt"
	"net/http"
	"time"
)

func HelloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	currentTime := time.Now()
	fmt.Fprintln(writer, "Hello World!", currentTime.Format("2006.01.02 15:04:05"))
}
