package handler

import (
	"connectDB/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllTodos(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Todo" : repository.FindAllTodos(c),
	})
}
func AddTodo()  {

}
func EditTodo()  {

}
func DeleteTodo()  {

}
