package main

import (
	"flag"

	"lab_3/app"
	"lab_3/helper/config"
)

func main() {
	env := flag.String("env", "local", "application runtime environment")
	flag.Parse()
	config.Loads("config.yml")
	config.SetEnv(*env)
	app.Rouer()
}
