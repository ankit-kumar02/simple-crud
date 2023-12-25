package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ankit-kumar02/simple-crud/model"
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
	fmt.Println(user)
	// User found, return the user details
	c.JSON(http.StatusOK, user)

}
