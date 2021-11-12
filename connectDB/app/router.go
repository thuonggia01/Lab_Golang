package app

import (
	"connectDB/handler"

	"github.com/gin-gonic/gin"
)

func router(app *gin.Engine) {
	todohadler := handler.NewTodoHandler()

	todo := app.Group("/todo")
	{
		todo.GET("/", todohadler.Index)
	}
}
