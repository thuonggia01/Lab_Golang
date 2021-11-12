package repoimlp

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"lab_3/model"
	repo "lab_3/repository"
)

type repositoryTodo struct{}

func NewRepositoryTodo() repo.RepositoryTodo {
	return &repositoryTodo{}
}

func (d *repositoryTodo) InsertTodo(to *model.Todo, db *gorm.DB) error {
	result := db.Create(&to)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Xóa Dữ Liệu By Id
func (d *repositoryTodo) DeleteByID(ID int64, db *gorm.DB) error {
	result := db.Delete(&model.Todo{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Update by ID

func (d *repositoryTodo) UpdateByID(todo *model.Todo, ID int64, db *gorm.DB) error {
	result := db.Exec("UPDATE todos SET description = ?, title = ? ,status = ? Where ID = ?", todo.Description, todo.Title, todo.Status, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Lấy tất cả dữ liệu có trong bảng
func (d *repositoryTodo) GetAllTodo(db *gorm.DB) ([]*model.Todo, error) {
	var todo []*model.Todo
	err := db.Find(&todo)
	if err.Error != nil {
		return nil, err.Error
	}

	return todo, nil
}

// Find By ID

func (d *repositoryTodo) FindTodoByID(id int64, db *gorm.DB) (*model.Todo, error) {
	todo := &model.Todo{}

	result := db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return todo, nil
}

func (d *repositoryTodo) GetTodoUsers(db *gorm.DB) (error, []model.TodoUsers) {
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

func (d *repositoryTodo) FindAllTodoUser(db *gorm.DB, title string, date string) ([]model.FindTodo, error) {
	a := []model.FindTodo{}
	result := db.Raw("SELECT  t.title ,u.name ,t.description, t.status ,t.created_at from public.todos as t INNER join public.users as u ON u.id = t.user_id where t.created_at >= '" + date + "'::date And t.created_at <= ('" + date + "'::date + '1 day'::interval) And t.title like '%" + title + "%'").Scan(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}
