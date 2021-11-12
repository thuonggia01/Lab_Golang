package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Adduser(c *gin.Context) {
	c.HTML(http.StatusOK, "addUser.html", gin.H{
		"title": "ADD USER",
	})
}
