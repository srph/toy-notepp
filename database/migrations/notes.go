package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigrateNotes() {
	database.Instance.MustExec(`
		CREATE TABLE notes(
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropNotes() {
	database.Instance.MustExec("DROP TABLE IF EXISTS notes")
}