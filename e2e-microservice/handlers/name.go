package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type NameHandler struct {
	logger *log.Logger
}

func NewNameHandler(logger *log.Logger) *NameHandler {
	return &NameHandler{logger: logger}
}

func (hello *NameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hello.logger.Printf("%s", data)
	fmt.Fprintf(w, "<h1>You asked for someone named: <small>%s</small></h1>", data)
}
