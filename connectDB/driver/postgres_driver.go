package driver

import (
	"connectDB/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect()(*gorm.DB, error){
db,err :=gorm.Open("postgres" ,"host=localhost port=5432 user=postgres dbname=demo  sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}

if err:=db.DB().Ping(); err != nil {
		panic(err)
	}
db.AutoMigrate(&model.Todo{})
	return db,nil
}
