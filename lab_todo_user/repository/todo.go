package repository

import (
	"github.com/jinzhu/gorm"

	"lab_3/model"
)

type RepositoryTodo interface {
	InsertTodo(to *model.Todo, db *gorm.DB) error
	DeleteByID(ID int64, db *gorm.DB) error
	UpdateByID(todo *model.Todo, ID int64, db *gorm.DB) error
	GetAllTodo(db *gorm.DB) ([]*model.Todo, error)
	FindTodoByID(id int64, db *gorm.DB) (*model.Todo, error)
	GetTodoUsers(db *gorm.DB) (error, []model.TodoUsers)
	FindAllTodoUser(db *gorm.DB, title string, date string) ([]model.FindTodo, error)
}
