package app

import (
	"github.com/gin-gonic/gin"

	"lab_3/handler"
	"lab_3/template"
)

func router(app *gin.Engine) {
	todohander := handler.NewTodotHandler()
	todotemplate := template.NewTodoTemplate()
	usetemplate := handler.NewuSerHandler()

	todo := app.Group("/")
	{
		todo.GET("", todotemplate.Alltodo)
		todo.GET(":ID", todohander.Delete)
		todo.POST(":ID", todohander.Update)
		todo.POST("add", todohander.Insert)
		todo.GET("add", todotemplate.Insert)
		todo.POST("find", todotemplate.FindTUsers)
		todo.GET("update/:ID", todotemplate.Update)
	}
	app.POST("/user", usetemplate.Insert)
	app.GET("/AddUser", template.Adduser)
	app.GET("/todoUser", todotemplate.TodoUsers)

}
