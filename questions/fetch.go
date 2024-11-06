package questions

import (
	"net/http"

	"github.com/anishNagula/pesuio-final-project/database"
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func FetchQuestion(c *gin.Context) {
	var request models.FetchQuestionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var question models.Question
	// Use GetDB() to access the database and preload the associated test cases
	if err := database.GetDB().Preload("TestCases").First(&question, request.QuestionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"question":  question.Question,
		"testCases": question.TestCases,
	})
}
