package controllers

import (
	"fmt"
	"net/http"
)

func (server *Server) TestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server Working")
}
