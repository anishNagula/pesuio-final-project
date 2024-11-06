package compiler

import (
	"os"
	"os/exec"

	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	var request models.RunRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	filename := "code.py"
	f, err := os.Create(filename)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create file"})
		return
	}
	defer f.Close()

	_, err = f.WriteString(request.Code)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not write to file"})
		return
	}

	cmd := exec.Command("python", filename, request.Input)
	output, err := cmd.CombinedOutput()

	if err != nil {
		c.JSON(400, gin.H{"error": "failed to run the code", "details": err.Error()})
		return
	}
	c.JSON(200, gin.H{"output": string(output)})
}
