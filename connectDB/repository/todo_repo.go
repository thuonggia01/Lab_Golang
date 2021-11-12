package repository

import (
	"connectDB/model"

	"github.com/jinzhu/gorm"
)

// func FindAllTodos() []model.Todo {
// 	db, err := driver.Connect()
// 	if err != nil {
// 		panic(err)
// 	}
// 	//defer db.Close()
// 	var todos []model.Todo
// 	db.Find(&todos)
// 	return todos
// }
// func NewTodo(newTodo model.Todo) {
// 	db, err := driver.Connect()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	//defer db.Close()
// 	db.Create(&newTodo)
// 	log.Println("created")
// }
// func UpdateTodo(newTodo model.Todo, idTodo int) {
// 	db, err := driver.Connect()
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.Model(&newTodo).Where("id = ?", idTodo).Update(newTodo)
// }
// func DeleteTodo(idTodo int) {
// 	db, err := driver.Connect()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.Delete(&model.Todo{}, idTodo)
// }

type TodoRepository interface {
	FindAllTodos(db *gorm.DB) ([]model.Todo, error)
	FindByID(db *gorm.DB, idTodo int) (*model.Todo, error)
	RemoveTodoById(db *gorm.DB, idTodo int) error
	Create(db *gorm.DB, todo *model.Todo) error
	Save(db *gorm.DB, todo *model.Todo, Id int) error
	GetTodoUsers(db *gorm.DB) (error, []model.TodoUsers)
	FindAllTodoUser(db *gorm.DB, title string, date string) ([]model.FindTodo, error)
}
