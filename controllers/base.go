package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"ecommerce/entity"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	server.DB.Debug().AutoMigrate(&entity.User{}, &entity.Cart{}, &entity.Product{})
	server.Router = mux.NewRouter()
	server.setupRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
