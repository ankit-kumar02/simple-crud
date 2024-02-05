package routes

import (
	"github.com/ankit-kumar02/simple-crud/controller"
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/repository"
	"github.com/ankit-kumar02/simple-crud/service"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UrlGeneration() {

	Router = gin.Default()
	Router.GET("/userdetails/:id", controller.UserDetails)
	Router.POST("/userdetails", controller.SaveUser)
	Router.DELETE("/userdetails/:id", controller.DeleteUser)
	Router.POST("/updateuserdetails/:id", controller.UpdateUser)
	DB, _ := model.InitDB()
	router := Router.Group("/api")
	bureauRouter := router.Group("/bureau")
	bureauRepository := repository.NewTagsRepositoryImpl(DB)
	bureauService := service.NewBureauServiceImple(bureauRepository)
	bureauController := controller.NewBureauController(bureauService)
	// bureauController := controller.BureauController{}
	bureauRouter.GET("", bureauController.FindAll)
	bureauRouter.GET("/:bureauId", bureauController.FindById)
	bureauRouter.POST("", bureauController.Create)
	bureauRouter.PATCH("/:bureauId", bureauController.Update)
	bureauRouter.DELETE("/:bureauId", bureauController.Delete)
}
