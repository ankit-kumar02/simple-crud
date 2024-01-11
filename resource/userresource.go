package resource

import (
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/gin-gonic/gin"
)

func FilterUserResponse(user *model.User) gin.H {
	if user == nil {
		return gin.H{}
	}

	filteredUser := gin.H{
		"ID":       user.ID,
		"Name":     user.Name,
		"Email":    user.Email,
		"Age":      user.Age,
		"Birthday": user.Birthday.String(),
		"Address":  user.Address,
	}

	return filteredUser
}
