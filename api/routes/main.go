package routes

import (
	"github.com/ThiagoMaylon/QuizAPI/api/controllers"
	"github.com/gin-gonic/gin"
)


func Routes(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/")
	{
		v1.GET("", controllers.GetAll)
		v1.GET("/topics/:id", controllers.GetTopicId)
	}
	return v1
}
