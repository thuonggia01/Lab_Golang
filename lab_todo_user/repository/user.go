package repository

import (
	"github.com/jinzhu/gorm"

	"lab_3/model"
)

type RepositoryUser interface {
	GetAllUsers(db *gorm.DB) ([]*model.User, error)
	InsertUsers(to model.User, db *gorm.DB) error
}
