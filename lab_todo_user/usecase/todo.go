package usecase

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"lab_3/helper/database"
	"lab_3/model"
	"lab_3/repository"
	repoimlp "lab_3/repository/repoImlp"
)

type UserCaseTodo interface {
	InsertTodo(to *model.Todo) error
	DeleteByID(ID int64) error
	UpdateByID(todo *model.Todo, ID int64) error
	GetAllTodo() ([]*model.Todo, error)
	GetTodoUsers() (error, []model.TodoUsers)
	FindTodoByID(id int64) (*model.Todo, error)
	TodoUserByTitleDate(title string, date string) ([]model.FindTodo, error)
}

type (
	userCaseTodo struct {
		db             *gorm.DB
		Todorepository repository.RepositoryTodo
	}
)

func NewuserCaseTodo() UserCaseTodo {
	return &userCaseTodo{
		db:             database.DB(),
		Todorepository: repoimlp.NewRepositoryTodo(),
	}
}

func (u *userCaseTodo) GetAllTodo() ([]*model.Todo, error) {
	list, err := u.Todorepository.GetAllTodo(u.db)
	return list, err
}

func (t *userCaseTodo) TodoUserByTitleDate(title string, date string) ([]model.FindTodo, error) {

	ftodo, err := t.Todorepository.FindAllTodoUser(t.db, title, date)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ftodo)

	return ftodo, err
}

func (u *userCaseTodo) GetTodoUsers() (error, []model.TodoUsers) {
	err, list := u.Todorepository.GetTodoUsers(u.db)
	return err, list
}

func (u *userCaseTodo) InsertTodo(to *model.Todo) error {
	err := u.Todorepository.InsertTodo(to, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *userCaseTodo) DeleteByID(ID int64) error {
	err := u.Todorepository.DeleteByID(ID, u.db)
	if err != nil {
		return err
	}
	return nil
}

// FindTodo
func (u *userCaseTodo) FindTodoByID(id int64) (*model.Todo, error) {
	todo, err := u.Todorepository.FindTodoByID(id, u.db)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func converStrToBool(status string) bool {
	if status == "" {
		return false
	}
	return true
}

func (u *userCaseTodo) UpdateByID(todo *model.Todo, ID int64) error {
	err := u.Todorepository.UpdateByID(todo, ID, u.db)
	if err != nil {
		return err
	}
	return nil
}
