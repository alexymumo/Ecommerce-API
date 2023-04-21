package controllers

import "ecommerce/api/middlewares"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/v1/test", middlewares.SetMiddlewareJson(s.TestServer)).Methods("GET")
	s.Router.HandleFunc("/v1/create", middlewares.SetMiddlewareJson(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/v1/delete", middlewares.SetMiddlewareJson(s.DeleteProductById)).Methods("DELETE")
	s.Router.HandleFunc("/v1/products", middlewares.SetMiddlewareJson(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/v1/update", middlewares.SetMiddlewareJson(s.UpdateProduct)).Methods("PUT")
}
