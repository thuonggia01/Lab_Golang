package main

import (
	"connectDB/driver"
	"connectDB/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	driver.Connect()
	router := gin.Default()
	router.Static("assets","./assets")
	router.LoadHTMLGlob("./assets/**/*.html")
	router.GET("/",handler.FindAllTodos)
	router.GET("/add", handler.AddTodo)
	//router.POST("/add", handler.AddTodo)
	router.GET("/edit/:id",handler.EditTodo)
	router.GET("/delete/:id",handler.DeleteTodo)
	router.Run(":8080")
}

