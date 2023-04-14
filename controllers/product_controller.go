package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ecommerce/entity"
)

var product []entity.Product

type controller struct{}

type ProductController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProductById(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

func (*controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prod, _ := ioutil.ReadAll(r.Body)
	var createProduct entity.Product
	prod = append(prod, prod...)
	json.NewEncoder(w).Encode(createProduct)
}

func (*controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	product = []entity.Product{
		{ID: 12, ProductName: "Mens Jacket", ProductImage: "jacket.jpg", Price: 23.9, Category: "Men", Description: "Man Jacket"},
		{ID: 123, ProductName: "Phone", ProductImage: "jacket.png", Price: 23.4, Category: "Tech", Description: "Electronics"},
	}
	json.NewEncoder(w).Encode(product)

}

func (*controller) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Product deleted")
}

func (*controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Updated product")
}
