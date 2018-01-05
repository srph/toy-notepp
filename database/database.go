package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Instance *sqlx.DB = sqlx.MustConnect("mysql", "root@tcp(localhost)/" + "failbook")