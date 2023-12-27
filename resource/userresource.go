package resource

import (
	"github.com/ankit-kumar02/simple-crud/model"
	"github.com/gin-gonic/gin"
)

func FilterUserResponse(user *model.User) gin.H {
	if user == nil {
		return gin.H{}
	}

	// Create a map or struct with only the fields you want to expose
	filteredUser := gin.H{
		"ID":           user.ID,
		"Name":         user.Name,
		"Age":          user.Age,
		"Birthday":     user.Birthday.String(), // Assuming Birthday is a time.Time field
		"MemberNumber": user.MemberNumber.String,
		// Add other fields as needed...
	}

	return filteredUser
}
