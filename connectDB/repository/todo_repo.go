package repository

import (
	"connectDB/driver"
	"connectDB/model"
	"log"
)

func FindAllTodos() ([]model.Todo){
db,err := driver.Connect()
	if err != nil {
		panic(err)
	}
//defer db.Close()
	var todos []model.Todo
	db.Find(&todos)
	return todos
}
func NewTodo (newTodo model.Todo){
	db,err := driver.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	//defer db.Close()
	db.Create(&newTodo)
	log.Println("created")
}
func UpdateTodo(newTodo model.Todo,idTodo int) {
	db,err := driver.Connect()
	if err != nil {
		panic(err)
	}
	db.Model(&newTodo).Where("id = ?", idTodo).Update(newTodo)
}
func DeleteTodo(idTodo int) {
	db,err := driver.Connect()
	if err != nil {
		panic(err)
	}

	db.Delete(&model.Todo{},idTodo)
}
