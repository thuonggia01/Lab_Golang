package repoimlp

import (
	"github.com/jinzhu/gorm"

	"lab_3/model"
	repo "lab_3/repository"
)

type repositoryUser struct{}

func NewRepositoryUser() repo.RepositoryUser {
	return &repositoryUser{}
}

func (d *repositoryUser) GetAllUsers(db *gorm.DB) ([]*model.User, error) {
	var Users []*model.User
	result := db.Find(&Users)
	if result.Error != nil {
		return nil, result.Error
	}
	return Users, nil
}

// Add
func (d *repositoryUser) InsertUsers(to model.User, db *gorm.DB) error {
	result := db.Save(&to)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
