package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThiagoMaylon/QuizAPI/api/quiz"
	"github.com/gin-gonic/gin"
)

// Middleware para configurar o cabeçalho Access-Control-Allow-Origin
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetAll(ctx *gin.Context) {
	quiz, err := quiz.GetQuiz("quiz.txt")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Houve um Erro"})
		return
	}
	ctx.JSON(http.StatusOK, quiz)
}

func GetTopicId(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	q, err := quiz.GetQuiz("quiz.txt")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Houve um erro ao obter os tópicos"})
		return
	}

	var topic quiz.Topics
	for _, t := range q.Topics {
		if t.Id == id {
			topic = t
			break
		}
	}


	ctx.JSON(http.StatusOK, topic)
}