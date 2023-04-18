package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ecommerce/api/entity"
	"ecommerce/api/responses"
)

/*
type ProductController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProductById(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}
*/

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	// deprecated in go 1.16
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product := entity.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	product.Prepare()
	err = product.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	products, err := product.SaveProduct(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	responses.JSON(w, http.StatusCreated, products)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	product := entity.Product{}
	products, err := product.GetProducts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)

}

func (server *Server) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Product deleted")
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Updated product")
}
