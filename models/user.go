package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

type User struct {
	ID int
	Username string
	Password string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts []Post
}