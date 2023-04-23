package controllers

import (
	"ecommerce/api/entity"
	"ecommerce/api/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := entity.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	users, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	//w.Header().Set()
	responses.JSON(w, http.StatusCreated, users)
}

func (server *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	users, err := user.GetUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}

func (server *Server) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := entity.User{}
	userid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteUser(server.DB, int32(userid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
