package main

import (
	"fmt"

	"ecommerce/controllers"
)

var server = controllers.Server{}

func init() {
	fmt.Print("Error occurred")
}

func Run() {
	server.Initialize("wew", "w2", "adf", "wewe", "wew", "sad")

}
