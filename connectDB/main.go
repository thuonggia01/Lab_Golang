package main

import (
	"connectDB/driver"
	"connectDB/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	driver.Connect()
	router := gin.Default()
	router.Static("assets", "./assets")
	router.LoadHTMLGlob("./assets/**/*.html")
	router.GET("/", handler.FindAllTodos)
	router.GET("/addTodo", handler.AddTodo)
	router.POST("/saveTodo", handler.SaveTodo)
	// router.GET("/edit/:id", handler.EditTodo)
	// router.POST("/update/:id", handler.UpdateTodo)
	router.GET("/delete/:id", handler.DeleteTodo)

	//task
	router.GET("/addTask", handler.AddUsersTask)
	// router.POST("/saveTask", handler.)
	router.Run(":8080")
}
