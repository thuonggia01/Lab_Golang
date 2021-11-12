package database

import (
	"connectDB/helpers/config"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	conn *gorm.DB
	once sync.Once
)

func DB(ctx *gin.Context) *gorm.DB {
	return Connection(ctx)
}

func Connection(ctx *gin.Context) *gorm.DB {
	connectTimeOut := 9
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s connect_timeout=%d",
		"localhost",
		"postgres",
		"demo",
		"disable",
		"postgres",
		connectTimeOut)
	once.Do(func() {
		conn = newConnection(ctx, connectionString)
	})
	return conn
}

func newConnection(ctx *gin.Context, connStr string) (db *gorm.DB) {
	env := config.GetEnvValue()
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	maxOpenConnections := env.Database.MaxOpenConnections
	if maxOpenConnections < runtime.GOMAXPROCS(0)*4 {
		maxOpenConnections = runtime.GOMAXPROCS(0) * 4
	}
	db.DB()
	db.DB().SetConnMaxLifetime(time.Hour * 24)
	db.DB().SetMaxIdleConns(env.Database.MaxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)

	switch config.GetEnv() {
	case config.EnvLocal:
		db.LogMode(true)
	case config.EnvTesting:
		db.LogMode(false)
	case config.EnvProduction:
		db.LogMode(false)
	}

	return db
}
