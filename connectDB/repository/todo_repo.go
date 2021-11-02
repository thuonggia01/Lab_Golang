package repository

import (
	"connectDB/driver"
	"connectDB/model"
	"github.com/gin-gonic/gin"
)

func FindAllTodos(c *gin.Context) ([]model.Todo){
db,_ := driver.Connect()
	var todos []model.Todo
	db.Find(&todos)

	return todos
}
//func SaveTodo(c *gin.Context) {
//	db,err := driver.Connect()
//	if err != nil {
//		panic(err)
//	}
//	newTodo :=model.Todo{
//		Model:       gorm.Model{},
//		Title:       c.PostForm(""),
//		Description: c.PostForm(""),
//		Status:      true,
//	}
//	db.Save(&newTodo)
//}
//func UpdateTodo(c *gin.Context) {
//	db,err := driver.Connect()
//	if err != nil {
//		panic(err)
//	}
//	newTodo :=model.Todo{
//		Model:       gorm.Model{},
//		Title:       c.PostForm(""),
//		Description: c.PostForm(""),
//		Status:      true,
//	}
//	db.Model(&newTodo).Where("id = ?", newTodo.ID).Update(newTodo)
//}
//func DeleteTodo(c *gin.Context) {
//	db,err := driver.Connect()
//	if err != nil {
//		panic(err)
//	}
//	newTodo :=model.Todo{
//		Model:       gorm.Model{},
//		Title:       c.PostForm(""),
//		Description: c.PostForm(""),
//		Status:      true,
//	}
//	db.Delete(&model.Todo{},newTodo.ID)
//}
