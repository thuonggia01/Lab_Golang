package app

import (
	"github.com/gin-gonic/gin"

	"lab_3/viewhelper"
)

func Rouer() error {
	app := setup()
	return app.Run()
}

func setup() *gin.Engine {
	app := gin.Default()
	app.Static("/assets", "./assets")
	app.SetFuncMap(viewhelper.Funcs())
	// app.LoadHTMLGlob("./template/*.html")
	app.LoadHTMLGlob("./template/**/*.html")
	router(app)
	return app
}
