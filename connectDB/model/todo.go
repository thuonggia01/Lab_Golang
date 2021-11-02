package model

import (
	"github.com/jinzhu/gorm"
)


type Todo struct {
	gorm.Model
	Title string
	Description string
	Status bool
}
