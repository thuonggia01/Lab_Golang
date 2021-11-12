package repository

import (
	"connectDB/model"

	"github.com/jinzhu/gorm"
)

var _ UserRepository = (*userRepository)(nil)

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
func (r *userRepository) FindAllUser(db *gorm.DB) ([]*model.User, error) {
	var Users []*model.User
	result := db.Find(&Users)
	if result.Error != nil {
		return nil, result.Error
	}
	return Users, nil
}

func (r *userRepository) Create(db *gorm.DB, user model.User) error {
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func (r *userRepository) FindByID(db *gorm.DB, idUser string) (*model.User, error) {
// 	var user model.User
// 	if err := db.Where("id = ?", idUser).Find(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) removeUserById(db *gorm.DB, idUser string) error {
// 	return db.Delete(&model.User{}, idUser).Error
// }

// func (r *userRepository) Save(db *gorm.DB, user *model.User) error {
// 	return db.Save(user).Error
// }
