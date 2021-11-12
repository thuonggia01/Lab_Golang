package repository

import (
	"connectDB/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindAllUser(db *gorm.DB) ([]*model.User, error)
	Create(db *gorm.DB, user model.User) error

	// FindByID(db *gorm.DB, idUser string) (*model.User, error)
	// removeUserById(db *gorm.DB, idUser string) error
	// Save(db *gorm.DB, user *model.User) error
}
