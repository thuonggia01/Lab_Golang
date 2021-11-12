package template

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"lab_3/model"
	"lab_3/usecase"
)

type TodoTemplateInterface interface {
	Alltodo(c *gin.Context)
	Update(c *gin.Context)
	Insert(c *gin.Context)
	TodoUsers(c *gin.Context)
	FindTUsers(c *gin.Context)
}

type data struct {
	Todo *model.Todo
	Time string
}

type todotTemplate struct {
	usecaseTodo usecase.UserCaseTodo
	usecaseUser usecase.UserCaseUser
}

func NewTodoTemplate() TodoTemplateInterface {
	return &todotTemplate{
		usecaseTodo: usecase.NewuserCaseTodo(),
		usecaseUser: usecase.NewUserCase(),
	}
}

func (t *todotTemplate) Alltodo(c *gin.Context) {
	data := t.formatData()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Data": data,
	})
}

func (t *todotTemplate) Insert(c *gin.Context) {
	user, err := t.usecaseUser.GetAllUsers()
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "ADD", "User": user,
	})
}

func (t *todotTemplate) Update(c *gin.Context) {
	ID := c.Param("ID")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		log.Println(err)
	}
	todo, err := t.usecaseTodo.FindTodoByID(id)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "update.html", gin.H{
		"Data": todo,
	})
}

//
// // List Todo User
func (t *todotTemplate) TodoUsers(c *gin.Context) {
	err, UseTodo := t.usecaseTodo.GetTodoUsers()
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "todoUser.html", gin.H{
		"title": "ADD USER",
		"User":  UseTodo,
	})
}

//List Find Todo User

func (t *todotTemplate) FindTUsers(c *gin.Context) {

	title := c.PostForm("findtitle")
	date := c.DefaultPostForm("finddate", "2022-12-5")

	fUseTodo, err := t.usecaseTodo.TodoUserByTitleDate(title, date)
	if err != nil {
		log.Println(err)
	}
	c.HTML(http.StatusOK, "findTodo.html", gin.H{
		"title": "Find",
		"ftodo": fUseTodo,
	})
}

func (t *todotTemplate) formatData() []*data {
	Todos, err := t.usecaseTodo.GetAllTodo()
	Data := make([]*data, 0, 1)
	if err != nil {
		log.Println("formatData", err)
	}
	for _, todo := range Todos {
		t := todo
		time := getTime(t)
		d := data{Todo: t, Time: time}
		Data = append(Data, &d)
	}
	return Data
}

func getTime(t *model.Todo) string {
	stime1C := t.Model.CreatedAt
	stime2C := t.Model.UpdatedAt
	var time time.Time = stime1C

	if !stime1C.Equal(stime2C) {
		time = stime2C
	}

	fmt.Print(!stime1C.Equal(stime2C))
	timefomat := time.Format("02/01/2006")
	return timefomat
}
