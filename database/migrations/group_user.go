package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigrateGroupUser() {
	database.Instance.MustExec(`
		CREATE TABLE group_user(
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL,
			group_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropGroupUser() {
	database.Instance.MustExec("DROP TABLE IF EXISTS group_user")
}