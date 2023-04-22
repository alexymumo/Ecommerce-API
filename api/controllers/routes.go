package controllers

import "ecommerce/api/middlewares"

func (s *Server) setupRoutes() {

	// product endpoints
	s.Router.HandleFunc("/v1/product/test", middlewares.SetMiddlewareJson(s.TestServer)).Methods("GET")
	s.Router.HandleFunc("/v1/product/create", middlewares.SetMiddlewareJson(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/v1/product/delete", middlewares.SetMiddlewareJson(s.DeleteProductById)).Methods("DELETE")
	s.Router.HandleFunc("/v1/product/products", middlewares.SetMiddlewareJson(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/v1/product/update", middlewares.SetMiddlewareJson(s.UpdateProduct)).Methods("PUT")

	// user end points
	s.Router.HandleFunc("/v1/users/register", middlewares.SetMiddlewareJson(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/v1/users/delete{id}", middlewares.SetMiddlewareJson(s.DeleteUserById)).Methods("DELETE")

}
