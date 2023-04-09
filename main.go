package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID           int     `json:"id"`
	ProductName  string  `json:"productname"`
	ProductImage string  `json:"productimage"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
}

var product []Product

/*

ID           int     `json:"id"`
	ProductName  string  `json:"productname"`
	ProductImage string  `json:"productimage"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
*/

func getProducts(w http.ResponseWriter, r *http.Request) {
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

func main() {
	product = append(product, Product{ID: 12, ProductName: "Mens Jacket", ProductImage: "jacket.jpg", Price: 23.9, Category: "Men", Description: "Man Jacket"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/products", getProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", getProductByID).Methods("GET")
	router.HandleFunc("/api/products", createProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", deleteProductByID).Methods("DELETE")
	router.HandleFunc("/api/products/{id}", updateProductByID).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
