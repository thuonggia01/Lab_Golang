package usecase

import (
	"github.com/jinzhu/gorm"

	"lab_3/helper/database"
	"lab_3/model"
	"lab_3/repository"
	repoimlp "lab_3/repository/repoImlp"
)

type UserCaseUser interface {
	Insert(User model.User) error
	GetAllUsers() ([]*model.User, error)
}

type usercase struct {
	db       *gorm.DB
	repoUser repository.RepositoryUser
}

func NewUserCase() UserCaseUser {
	return &usercase{
		db:       database.DB(),
		repoUser: repoimlp.NewRepositoryUser(),
	}
}

func (u *usercase) Insert(User model.User) error {
	err := u.repoUser.InsertUsers(User, u.db)
	if err != nil {
		return err
	}
	return nil
}

func (u *usercase) GetAllUsers() ([]*model.User, error) {
	uses, err := u.repoUser.GetAllUsers(u.db)
	if err != nil {
		return nil, err
	}
	return uses, nil
}
