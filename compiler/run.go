package compiler

import (
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	var request models.RunRequest
	c.BindJSON(&request)
	// this runs the code and returns the stdout and stderr

	// implement

	c.JSON(200, gin.H{
		"output": "Hello World",
	})
}
