package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Id         string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     time.Time `json:"created"`
}
type listTodo []Todo

var Todos = listTodo{
	{
		Id:          "1",
		Title:       "Golang",
		Description: "learn",
		Created:     time.Now(),
	},
	{
		Id:          "2",
		Title:       "Golang",
		Description: "learn",
		Created:     time.Now(),
	},
}

//getAllTodos
func getTodos(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, Todos)
	part , err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(part)
	}
	err1 := part.Execute(c.Writer,Todos)
	if err1 != nil {
		log.Println(err1)
	}
}

//getByTitle
func getById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range Todos {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// create Todos
func postTodos(c *gin.Context) {
	c.PostFormMap()
	id := c.Post
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		log.Fatal(err)
	}
	Todos = append(Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
	c.Redirect(http.StatusFound, "/todos")
}

//update Todos
func updateTodos(c *gin.Context) {
	id := c.Param("id")
	var updateTodo Todo
	if err := c.BindJSON(&updateTodo); err != nil {
		return
	}
	for i, oldTodo := range Todos {
		if oldTodo.Id == id {
			oldTodo.Title = updateTodo.Title
			oldTodo.Description = updateTodo.Description
			Todos = append(Todos[:i], oldTodo)
			c.IndentedJSON(http.StatusCreated, oldTodo)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/todos")
}

//delete Todos
func deleteTodos(c *gin.Context) {

	id := c.Param("id")
	for i, todo := range Todos {
		if todo.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/todos")
}
//save Todos
func addTodo(c *gin.Context)  {
	part , err := template.ParseFiles("./templates/add.html")
	if err != nil {
		log.Println(part)
	}
	err1 := part.Execute(c.Writer,"hello")
	if err1 != nil {
		log.Println(err1)
	}
}
//edit Todos
func editTodo(c *gin.Context)  {
	var id = c.Param("Id")
	part , err := template.ParseFiles("./templates/edit.html")
	if err != nil {
		log.Println(part)
	}
	err1 := part.Execute(c.Writer,id)
	if err1 != nil {
		log.Println(err1)
	}
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getById)
	router.POST("/create/todos", postTodos)
	router.PATCH("/todos/update/:id", updateTodos)
	router.DELETE("/todos/delete/:id", deleteTodos)
	router.GET("/todos/edit/:id", editTodo)
	router.GET("/add", addTodo)
	router.Run("localhost:8080")
	////ghi file
	//jsonString, _ := json.Marshal(Todos)
	//ioutil.WriteFile("todos.json", jsonString, os.ModePerm)
	//
}
