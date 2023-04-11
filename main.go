package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)

	for _, productId := range product {
		if string(productId.ID) == key {
			json.NewEncoder(w).Encode(productId)
		}
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prod, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(prod))

	var product1 entity.Product
	json.Unmarshal(prod, &product1)
	product = append(product, product1)
	json.NewEncoder(w).Encode(product1)

}

func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//id := vars["id"]

}

func updateProductByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world")

}

func testCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to server")
}

func main() {
	//product = append(product, Product{})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/products", getProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", getProductByID).Methods("GET")
	router.HandleFunc("/api/product", createProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", deleteProductByID).Methods("DELETE")
	router.HandleFunc("/api/products/{id}", updateProductByID).Methods("PUT")
	router.HandleFunc("/api/test", testCode).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
