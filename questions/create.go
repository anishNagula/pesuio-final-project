package questions

import (
	"net/http"

	"github.com/anishNagula/pesuio-final-project/database"
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var request models.CreateQuestionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate input
	if request.Question == "" || len(request.TestCases) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Question and test cases are required"})
		return
	}

	// Create the question in the database using GetDB()
	question := models.Question{
		Question:  request.Question,
		TestCases: request.TestCases,
	}

	// Using GetDB() to access the db instance
	if err := database.GetDB().Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "questionID": question.ID})
}
