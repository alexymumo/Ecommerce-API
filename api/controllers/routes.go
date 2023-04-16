package controllers

import "ecommerce/api/middlewares"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/test", middlewares.SetMiddlewareJson(s.TestServer)).Methods("GET")
	s.Router.HandleFunc("/create", middlewares.SetMiddlewareJson(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/delete", middlewares.SetMiddlewareJson(s.DeleteProductById)).Methods("DELETE")
	s.Router.HandleFunc("/all", middlewares.SetMiddlewareJson(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/update", middlewares.SetMiddlewareJson(s.UpdateProduct)).Methods("PUT")
}
