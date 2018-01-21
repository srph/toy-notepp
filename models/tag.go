package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

type Tag struct {
	ID int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	Notes []Note `gorm:many2many:note_tag`
}