package auth

import (
	"github.com/anishNagula/pesuio-final-project/database"
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var request models.SignUpRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
	}

	err = database.CreateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "couldn't create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
