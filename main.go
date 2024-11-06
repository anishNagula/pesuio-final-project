package main

import (
	"log"

	"github.com/anishNagula/pesuio-final-project/auth"
	"github.com/anishNagula/pesuio-final-project/compiler"
	"github.com/anishNagula/pesuio-final-project/questions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth/signin", auth.Signin)
	router.POST("/auth/signup", auth.Signup)

	router.POST("/run", compiler.Run)

	router.POST("/question/create", questions.CreateQuestion)
	router.POST("/question/fetch", questions.FetchQuestion)
	log.Println("Server is running on port 6969!")
	router.Run(":6969")
}
