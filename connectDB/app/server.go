package app

import "github.com/gin-gonic/gin"

// Startup the server
func Route() error {
	app := setup()
	return app.Run()
}

func setup() *gin.Engine {
	app := gin.New()
	app.Static("/assets", "./assets")
	//app.SetFuncMap(viewhelper.Funcs())
	app.LoadHTMLGlob("templates/*.html")
	router(app)
	return app
}
