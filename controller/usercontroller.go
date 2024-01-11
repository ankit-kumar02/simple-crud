package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/ankit-kumar02/simple-crud/resource"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB, err = model.InitDB()

func UserDetails(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := model.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	filteredUser := resource.FilterUserResponse(&user)

	c.JSON(http.StatusOK, gin.H{"data": filteredUser})

}

func SaveUser(c *gin.Context) {
	var request struct {
		Name     string
		Email    string
		Age      int
		Birthday string
		Address  string
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", request.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Age:      request.Age,
		Birthday: &parsedTime,
		Address:  request.Address,
	}

	model.DB.Create(&user)
	filteredUser := resource.FilterUserResponse(&user)

	c.JSON(http.StatusOK, gin.H{"data": filteredUser})
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := model.DB.Delete(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User Deleted Succesfully"})

}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	var request map[string]interface{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := model.DB.Model(&user).Updates(request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	filteredUser := resource.FilterUserResponse(&user)

	c.JSON(http.StatusOK, gin.H{"data": filteredUser})
}
