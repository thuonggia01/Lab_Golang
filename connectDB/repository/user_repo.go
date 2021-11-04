package repository

import (
	"connectDB/driver"
	"connectDB/model"
)

func FindAllUsers() []model.User {
	db, err := driver.Connect()
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	var users []model.User
	db.Find(&users)
	return users
}
