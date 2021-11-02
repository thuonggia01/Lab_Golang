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
	//router.GET("/add", handler.AddTodo)
	router.GET("/",handler.FindAllTodos)
	//router.GET("/detail/:name", handler.DeleteTodo)
	router.Run(":8080")
}

