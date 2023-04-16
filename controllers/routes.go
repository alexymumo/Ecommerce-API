package controllers

import "ecommerce/middlewares"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.TestServer)).Methods("GET")
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.DeleteProductById)).Methods("DELETE")
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.UpdateProduct)).Methods("PUT")
}
