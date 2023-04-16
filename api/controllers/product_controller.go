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

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "All Products")
	/*product = []entity.Product{
		{ID: 12, ProductName: "Mens Jacket", ProductImage: "jacket.jpg", Price: 23.9, Category: "Men", Description: "Man Jacket"},
		{ID: 123, ProductName: "Phone", ProductImage: "jacket.png", Price: 23.4, Category: "Tech", Description: "Electronics"},
	}

	json.NewEncoder(w).Encode(product)
	*/

}

func (server *Server) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Product deleted")
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Updated product")
}
