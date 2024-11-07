package auth

import (
	"github.com/anishNagula/pesuio-final-project/database"
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	var request models.SignInRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
	}

	success, err := database.CheckPassword(request.Username, request.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "server error: ",
		})
		return
	}

	if !success {
		c.JSON(401, gin.H{
			"error": "invalid username/password",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
