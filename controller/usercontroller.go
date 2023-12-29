package controller

import (
	"database/sql"
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

	// Query the database to find the user with the specified ID
	if err := model.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Other database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	filteredUser := resource.FilterUserResponse(&user)

	c.JSON(http.StatusOK, gin.H{"data": filteredUser})

}

func SaveUser(c *gin.Context) {
	var request struct {
		Name         string
		Email        *string
		Age          int
		Birthday     string
		MemberNumber string
		ActivatedAt  time.Time
		// other fields...
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the string into a time.Time value
	parsedTime, err := time.Parse("2006-01-02 15:04:05", request.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	// Create user
	user := model.User{
		Name:         request.Name,
		Email:        request.Email,
		Age:          uint8(request.Age),
		Birthday:     &parsedTime,
		MemberNumber: sql.NullString{String: request.MemberNumber, Valid: request.MemberNumber != ""},
	}

	model.DB.Create(&user)

	// Filter the user response before sending it to the client
	filteredUser := resource.FilterUserResponse(&user)

	c.JSON(http.StatusOK, gin.H{"data": filteredUser})
}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	// Query the database to find the user with the specified ID
	if err := model.DB.Delete(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Other database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User Deleted Succesfully"})

}
