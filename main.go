package main

import (
	"fmt"

	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/routes"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	viper.AutomaticEnv()

	model.InitDB()
	routes.UrlGeneration()
	routes.Router.Run("localhost:8082")
}
