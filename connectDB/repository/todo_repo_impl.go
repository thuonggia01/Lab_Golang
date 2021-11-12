package repository

import (
	"connectDB/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

var _ TodoRepository = (*todoRepository)(nil)

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}
func (r *todoRepository) FindAllTodos(db *gorm.DB) ([]model.Todo, error) {
	var todos []model.Todo
	err := db.Find(&todos)
	if err.Error != nil {
		return nil, err.Error
	}

	return todos, nil
}

func (r *todoRepository) FindByID(db *gorm.DB, idTodo int) (*model.Todo, error) {
	var todo model.Todo
	if err := db.Where("id = ?", idTodo).Find(&todo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) RemoveTodoById(db *gorm.DB, idTodo int) error {
	result := db.Delete(&model.Todo{}, idTodo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *todoRepository) Create(db *gorm.DB, todo *model.Todo) error {
	result := db.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *todoRepository) Save(db *gorm.DB, todo *model.Todo, Id int) error {
	result := db.Exec("UPDATE todos SET description = ?, title = ? ,status = ? Where ID = ?", todo.Description, todo.Title, todo.Status, Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *todoRepository) GetTodoUsers(db *gorm.DB) (error, []model.TodoUsers) {
	a := []model.TodoUsers{}
	result := db.Raw("SELECT  Title , Status , Name FROM public.todos INNER JOIN public.users ON todos.user_id = users.id WHERE todos.deleted_at is null").Scan(&a)
	for _, todo := range a {
		fmt.Println(todo)
	}

	if result.Error != nil {
		return result.Error, nil
	}
	return nil, a
}
func (r *todoRepository) FindAllTodoUser(db *gorm.DB, title string, date string) ([]model.FindTodo, error) {
	a := []model.FindTodo{}
	result := db.Raw("SELECT  t.title ,u.name ,t.description, t.status ,t.created_at from public.todos as t INNER join public.users as u ON u.id = t.user_id where t.created_at >= '" + date + "'::date And t.created_at <= ('" + date + "'::date + '1 day'::interval) And t.title like '%" + title + "%'").Scan(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}
