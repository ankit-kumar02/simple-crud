package routes

import (
	"github.com/ankit-kumar02/simple-crud/controller"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UrlGeneration() {
	Router = gin.Default()
	Router.GET("/userdetails/:id", controller.UserDetails)
	Router.POST("/userdetails", controller.SaveUser)
}
