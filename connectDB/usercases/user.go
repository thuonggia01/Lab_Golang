package usercases

import (
	"connectDB/helpers/database"
	"connectDB/model"
	"connectDB/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var _ UserUseCase = (*userUseCase)(nil)

type (
	UserUseCase interface {
		Create(c *gin.Context, user model.User) error
		FindAllUser(c *gin.Context) ([]*model.User, error)
	}

	userUseCase struct {
		db             *gorm.DB
		userRepository repository.UserRepository
	}
)

func NewUserUseCase(ctx *gin.Context) UserUseCase {
	return &userUseCase{
		db:             database.DB(ctx),
		userRepository: repository.NewUserRepository(),
	}
}
func (u *userUseCase) Create(c *gin.Context, user model.User) error {
	err := u.userRepository.Create(u.db, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) FindAllUser(c *gin.Context) ([]*model.User, error) {
	uses, err := u.userRepository.FindAllUser(u.db)
	if err != nil {
		return nil, err
	}
	return uses, nil
}
