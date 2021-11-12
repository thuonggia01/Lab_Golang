package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"lab_3/model"
	"lab_3/usecase"
)

type TodoInterface interface {
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Insert(c *gin.Context)
}

type todotHandler struct {
	usecase usecase.UserCaseTodo
}

func NewTodotHandler() TodoInterface {
	return &todotHandler{
		usecase: usecase.NewuserCaseTodo(),
	}
}

func (t *todotHandler) Insert(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	UserID := c.DefaultPostForm("preform", "client")

	userID, err1 := strconv.ParseInt(UserID, 10, 64)

	if err1 != nil {
		log.Println(err1)
	}

	todo := model.Todo{Title: title, Description: description, Status: false, UserID: userID}
	err := t.usecase.InsertTodo(&todo)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)

}

func (t *todotHandler) Update(c *gin.Context) {
	ID := c.Param("ID")
	title := c.PostForm("title")
	description := c.PostForm("description")
	tempStatus := c.Request.FormValue("status")
	id, err1 := strconv.ParseInt(ID, 10, 64)
	if err1 != nil {
		log.Println(err1)
	}
	Status := converStrToBool(tempStatus)
	todo := model.Todo{Title: title, Description: description, Status: Status}
	//
	err := t.usecase.UpdateByID(&todo, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
}

func (t *todotHandler) Delete(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println("Parse ID failed: ", err)
	}
	err1 := t.usecase.DeleteByID(id)
	if err1 != nil {
		log.Println(err1)
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
}

func converStrToBool(status string) bool {
	if status == "" {
		return false
	}
	return true
}
