package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("%s", data)
		fmt.Fprintf(w, "<h1>Hello I think you just called me: <small>%s</small></h1>", data)
	})
	http.ListenAndServe(":9090", nil)
}
