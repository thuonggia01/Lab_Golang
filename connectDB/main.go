package main

import (
	"connectDB/app"
	"connectDB/helpers/config"
	"flag"
	"os"
)

func main() {
	env := flag.String("env", "local", "application runtime environment")
	flag.Parse()
	config.Loads("config.yml")

	switch *env {
	case config.EnvProduction:
		config.SetEnv(config.EnvProduction)
	case config.EnvTesting:
		config.SetEnv(config.EnvTesting)
	case config.EnvLocal:
		config.SetEnv(config.EnvLocal)
	default:
		os.Exit(1)
	}

	err := app.Route()
	if err != nil {
		panic(err.Error())
	}
}
