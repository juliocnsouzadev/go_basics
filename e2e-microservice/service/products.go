package service

import (
	"e2e-microservice/models"
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

// GetProducts returns a list of products
func GetProducts() []*models.Product {
	return productList
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func ToJSON(products []*models.Product, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(products)
}

// productList is a hard coded list of products for this
// example data source
var productList = []*models.Product{
	&models.Product{
		ID:          uuid.New().String(),
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&models.Product{
		ID:          uuid.New().String(),
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
