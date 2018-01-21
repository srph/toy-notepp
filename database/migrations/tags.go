package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigrateTags() {
	database.Instance.MustExec(`
		CREATE TABLE tags(
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		)
	`)
}

func DropTags() {
	database.Instance.MustExec("DROP TABLE IF EXISTS tags")
}