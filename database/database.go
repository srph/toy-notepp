package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/srph/toy-notepp/lib/settings"
)

var Instance *sqlx.DB

func Init() {
	Instance = sqlx.MustConnect(
		"mysql",
		// "root@tcp(localhost)/failbook"
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			settings.DatabaseUser,
			settings.DatabasePassword,
			settings.DatabaseHost,
			settings.DatabaseName,
		),
	)
}