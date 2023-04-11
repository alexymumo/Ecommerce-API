package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommerce/entity"
)

var product []entity.Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	product = []entity.Product{
		{ID: 12, ProductName: "Mens Jacket", ProductImage: "jacket.jpg", Price: 23.9, Category: "Men", Description: "Man Jacket"},
		{ID: 123, ProductName: "Phone", ProductImage: "jacket.png", Price: 23.4, Category: "Tech", Description: "Electronics"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting product by id")
}

func createProduct(w http.ResponseWriter, r *http.Request) {

}

func deleteProductByID(w http.ResponseWriter, r *http.Request) {

}

func updateProductByID(w http.ResponseWriter, r *http.Request) {

}

func testCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to server")
}
