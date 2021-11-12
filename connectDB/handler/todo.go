package handler

import (
	"connectDB/usercases"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ HandlerTodo = (*handlerTodo)(nil)

type HandlerTodo interface {
	Index(c *gin.Context)
}
type (
	handlerTodo struct {
		todoUseCase usercases.TodoUseCase
	}
)

func NewTodoHandler() HandlerTodo {
	return &handlerTodo{}
}
func (h *handlerTodo) Index(c *gin.Context) {
	todos, err := usercases.NewTodoUseCase(c).GetAllTodos(c)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "template/index.html", gin.H{
		"todo": todos,
	})
}

// func AddTodo(c *gin.Context) {
// 	c.HTML(http.StatusOK, "add.html", gin.H{})
// }
// func SaveTodo(c *gin.Context) {
// 	var newTodo = model.Todo{
// 		Title:       c.PostForm("title"),
// 		Description: c.PostForm("description"),
// 		Status:      c.Request.PostForm.Has("status"),
// 	}
// 	log.Println(newTodo.Description, newTodo.ID, "**********************")
// 	repository.NewTodo(newTodo)
// 	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
// }
// func EditTodo(c *gin.Context) {
// 	c.HTML(http.StatusOK, "edit.html", gin.H{})
// }

// func UpdateTodo(c *gin.Context) {
// 	id := c.Param("id")
// 	i, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	newTodo := model.Todo{
// 		Model:       gorm.Model{},
// 		Title:       c.PostForm("title"),
// 		Description: c.PostForm("description"),
// 		Status:      true,
// 	}
// 	repository.UpdateTodo(newTodo, i)
// 	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
// }
// func DeleteTodo(c *gin.Context) {
// 	id := c.Param("id")
// 	i, err := strconv.Atoi(id)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	repository.DeleteTodo(i)
// 	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
// }
