package handlers

import (
	"e2e-microservice/service"
	"log"
	"net/http"
)

type ProudctsHandler struct {
	logger *log.Logger
}

func NewProudctsHandler(logger *log.Logger) *ProudctsHandler {
	return &ProudctsHandler{logger: logger}
}

func (products *ProudctsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := service.GetProducts()
		err := service.ToJSON(list, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		products.logger.Printf("POST not implemented")
	case http.MethodPut:
		products.logger.Printf("PUT not implemented")
		http.Error(w, "", http.StatusMethodNotAllowed)
	case http.MethodPatch:
		products.logger.Printf("PATCH not implemented")
		http.Error(w, "", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		products.logger.Printf("DELETE not implemented")
		http.Error(w, "", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}

}
