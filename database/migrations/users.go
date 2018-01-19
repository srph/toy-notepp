package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigrateUsers() {
	database.Instance.MustExec(`
		CREATE TABLE users(
			id INT NOT NULL AUTO_INCREMENT,
			username VARCHAR(100) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE (username)
		)
	`)
}

func DropUsers() {
	database.Instance.MustExec("DROP TABLE IF EXISTS users")
}