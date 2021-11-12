package usercases

import (
	"connectDB/helpers/database"
	"connectDB/model"
	"connectDB/repository"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var _ TodoUseCase = (*todoUseCase)(nil)

type (
	TodoUseCase interface {
		GetAllTodos(c *gin.Context) ([]model.Todo, error)
		GetByID(c *gin.Context, idTodo int) (*model.Todo, error)
		GetTodoUsers(c *gin.Context) (error, []model.TodoUsers)
		FindTodoUserByTitleDate(c *gin.Context, title string, date string) ([]model.FindTodo, error)
		DeleteTodoById(c *gin.Context, idTodo int) error
		Create(c *gin.Context, todo *model.Todo) error
		Save(c *gin.Context, todo *model.Todo, id int) error
	}

	todoUseCase struct {
		db             *gorm.DB
		todoRepository repository.TodoRepository
	}
)

func NewTodoUseCase(ctx *gin.Context) TodoUseCase {
	return &todoUseCase{
		db:             database.DB(ctx),
		todoRepository: repository.NewTodoRepository(),
	}
}
func (t *todoUseCase) GetAllTodos(c *gin.Context) ([]model.Todo, error) {
	todos, err := t.todoRepository.FindAllTodos(t.db)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (t *todoUseCase) GetByID(c *gin.Context, idTodo int) (*model.Todo, error) {

	todo, err := t.todoRepository.FindByID(t.db, idTodo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoUseCase) DeleteTodoById(c *gin.Context, idTodo int) error {
	err := t.todoRepository.RemoveTodoById(t.db, idTodo)
	if err != nil {
		return err
	}
	return nil
}
func (t *todoUseCase) Create(c *gin.Context, todo *model.Todo) error {
	err := t.todoRepository.Create(t.db, todo)
	if err != nil {
		return err
	}
	return nil
}
func (t *todoUseCase) Save(c *gin.Context, todo *model.Todo, id int) error {
	err := t.todoRepository.Save(t.db, todo, id)
	if err != nil {
		return err
	}
	return nil
}
func (t *todoUseCase) FindTodoUserByTitleDate(c *gin.Context, title string, date string) ([]model.FindTodo, error) {
	ftodo, err := t.todoRepository.FindAllTodoUser(t.db, title, date)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ftodo)

	return ftodo, err
}
func (t *todoUseCase) GetTodoUsers(c *gin.Context) (error, []model.TodoUsers) {
	err, list := t.todoRepository.GetTodoUsers(t.db)
	return err, list
}
func converStrToBool(status string) bool {
	if status == "" {
		return false
	}
	return true
}
