package main

import (
	"net/http"

	"github.com/ThiagoMaylon/QuizAPI/api/quiz"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/v1", func(ctx *gin.Context) {
		quiz, err := quiz.GetQuiz("quiz.txt")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Houve um Erro"})
			return
		}
		ctx.JSON(http.StatusOK, quiz)
	})

	r.Run()
}
