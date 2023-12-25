package main

import (
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/routes"
)

func main() {
	model.InitDB()
	routes.UrlGeneration()
	routes.Router.Run("localhost:8082")
}
