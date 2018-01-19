package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

type Post struct {
	ID int
	User *User
	UserID int
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}