package questions

import (
	"github.com/anishNagula/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func FetchQuestion(c *gin.Context) {
	var request models.FetchQuestionRequest
	c.BindJSON(&request)

	var question models.Question
	c.JSON(200, gin.H{
		"success":  true,
		"question": question,
	})

}
