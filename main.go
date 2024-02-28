package main

import (
	"github.com/ThiagoMaylon/QuizAPI/api/controllers"
	"github.com/ThiagoMaylon/QuizAPI/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(controllers.CORSMiddleware())
	routes.Routes(r)
	r.Run()
}
