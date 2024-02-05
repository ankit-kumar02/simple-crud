package controller

import (
	"net/http"

	"github.com/ankit-kumar02/simple-crud/helper"
	"github.com/ankit-kumar02/simple-crud/request"
	"github.com/ankit-kumar02/simple-crud/resource"
	"github.com/ankit-kumar02/simple-crud/service"
	"github.com/gin-gonic/gin"
)

type BureauController struct {
	bureauservice service.BureauService
}

func NewBureauController(service service.BureauService) *BureauController {
	return &BureauController{bureauservice: service}
}

func (controller *BureauController) Create(ctx *gin.Context) {
	createbureaurequest := request.CreateBureauRequest{}
	err := ctx.ShouldBindJSON(&createbureaurequest)
	helper.ErrorPanic(err)

	controller.bureauservice.Create(createbureaurequest)

	webResponse := resource.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BureauController) Update(ctx *gin.Context) {
	updatebureaurequest := request.UpdateBureauRequest{}
	err := ctx.ShouldBindJSON(&updatebureaurequest)
	helper.ErrorPanic(err)
	bureauId := ctx.Param("bureauId")

	helper.ErrorPanic(err)
	updatebureaurequest.BureauReferenceNumber = bureauId

	controller.bureauservice.Update(updatebureaurequest)
	webResponse := resource.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BureauController) Delete(ctx *gin.Context) {
	bureauId := ctx.Param("bureauId")
	// Id, err := strconv.Atoi(bureauId)
	// helper.ErrorPanic(err)
	controller.bureauservice.Delete(bureauId)

	webResponse := resource.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BureauController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("bureauId")
	// id, err := strconv.Atoi(tagId)
	// helper.ErrorPanic(err)

	bureauResponse := controller.bureauservice.FindById(tagId)

	webResponse := resource.Response{
		Code:   200,
		Status: "Ok",
		Data:   bureauResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BureauController) FindAll(ctx *gin.Context) {
	bureauResponse := controller.bureauservice.FindAll()

	webResponse := resource.Response{
		Code:   200,
		Status: "Ok",
		Data:   bureauResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
