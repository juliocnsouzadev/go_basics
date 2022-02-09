package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

func NewHelloHandler(logger *log.Logger) *HelloHandler {
	return &HelloHandler{logger: logger}
}

func (hello *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hello.logger.Printf("%s", data)
	fmt.Fprintf(w, "<h1>Hello: <small>%s</small></h1>", data)
}
