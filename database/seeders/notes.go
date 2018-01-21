package seeders

import (
	"github.com/srph/toy-notepp/database"
)

func SeedNotes() {
	database.Instance.MustExec(
		"TRUNCATE TABLE notes",
	)

	for i := 0; i < NOTE_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO notes(
				id,
				user_id,
				content
			) VALUES(
				?,
				?,
				?
			)`,
			id,
			(i % (USER_COUNT + 1)),
			"bla i'm a cool post",
		)
	}
}