package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Todo struct {
		gorm.Model
		UserID      int64
		Title       string
		Description string
		Status      bool
	}
	TodoUsers struct {
		Title  string
		Status bool
		Name   string
	}

	FindTodo struct {
		Title       string
		Name        string
		CreatedAt   time.Time
		Description string
		Status      bool
	}
)
