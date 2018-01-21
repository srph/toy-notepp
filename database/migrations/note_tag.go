package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigrateNoteTag() {
	database.Instance.MustExec(`
		CREATE TABLE note_tag(
			id INT NOT NULL AUTO_INCREMENT,
			tag_id INT NOT NULL,
			note_id INT NOT NULL,
			PRIMARY KEY (id)
		)
	`)
}

func DropNoteTag() {
	database.Instance.MustExec("DROP TABLE IF EXISTS note_tag")
}