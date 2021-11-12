package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"lab_3/model"
	"lab_3/usecase"
)

type usertInterface interface {
	Insert(c *gin.Context)
}

type usertHandler struct {
	usecase usecase.UserCaseUser
}

func NewuSerHandler() usertInterface {
	return &usertHandler{
		usecase: usecase.NewUserCase(),
	}
}

func (u *usertHandler) Insert(c *gin.Context) {
	name := c.DefaultPostForm("name", "client")
	user := model.User{Name: name}

	err := u.usecase.Insert(user)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
}
