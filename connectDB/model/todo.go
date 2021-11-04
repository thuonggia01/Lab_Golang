package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Todos []Todo //`gorm:" ForeignKey: UserId "`
}
type Todo struct {
	gorm.Model
	Title       string
	Description string
	Status      bool
	UserID      uint
}
