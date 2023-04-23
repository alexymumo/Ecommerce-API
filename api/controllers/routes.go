package controllers

import "ecommerce/api/middlewares"

func (s *Server) setupRoutes() {

	// test endpoint
	s.Router.HandleFunc("/v1/product/test", middlewares.SetMiddlewareJson(s.TestServer)).Methods("GET")

	// product endpoints
	s.Router.HandleFunc("/v1/product", middlewares.SetMiddlewareJson(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/v1/product{id}", middlewares.SetMiddlewareJson(s.DeleteProductById)).Methods("DELETE")
	s.Router.HandleFunc("/v1/products", middlewares.SetMiddlewareJson(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/v1/product/{id}", middlewares.SetMiddlewareJson(s.UpdateProduct)).Methods("PUT")

	// user endpoints
	s.Router.HandleFunc("/v1/user", middlewares.SetMiddlewareJson(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/v1/user/{id}", middlewares.SetMiddlewareJson(s.DeleteUserById)).Methods("DELETE")
	s.Router.HandleFunc("/v1/user", middlewares.SetMiddlewareJson(s.GetAllUsers)).Methods("GET")
}
