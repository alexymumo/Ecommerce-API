package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"ecommerce/api/entity"
	"ecommerce/api/responses"

	"github.com/gorilla/mux"
)

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

/*
Get all products
*/
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
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	product := entity.Product{}
	err = server.DB.Debug().Model(entity.Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Not found"))
		return
	}
	_, err = product.DeleteProduct(server.DB, uint64(id))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")

}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Updated product")
}

func (server *Server) SearchProducts(w http.ResponseWriter, r *http.Request) {

}
