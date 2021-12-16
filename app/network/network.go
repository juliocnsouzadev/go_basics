package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DoGet(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
