package controllers

import (
	"{{ .ProjectName }}/forms"
	"{{ .ProjectName }}/models"
	"github.com/gin-gonic/gin"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

var userModel = new(models.UserModel)

// UserController : struct for user controller
type UserController struct{}

// Create : function to create user 
func (user *UserController) Create(c *gin.Context) {
	var data forms.CreateUserCommand
	if c.BindJSON(&data) != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := userModel.Create(data)
	if err != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(406, gin.H{"message": "User could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

// Get : funcation to get user
func (user *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := userModel.Get(id)
	if err != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(404, gin.H{"message": "User not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

// Find : List all users
func (user *UserController) Find(c *gin.Context) {
	list, err := userModel.Find()
	if err != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

// Update : Function to update users based on ID
func (user *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	data := forms.UpdateUserCommand{}

	if c.BindJSON(&data) != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	err := userModel.Update(id, data)
	if err != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(406, gin.H{"message": "user count not be updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "User updated"})
}
// Delete : Controller function to delete user based on ID
func (user *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := userModel.Delete(id)
	if err != nil {
		{{ .Logging.Messages.Error }}
		c.JSON(406, gin.H{"message": "User could not be deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "User deleted"})
}
