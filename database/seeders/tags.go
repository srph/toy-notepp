package seeders

import (
	"github.com/srph/toy-notepp/database"
)

func SeedTags() {
	database.Instance.MustExec(
		"TRUNCATE TABLE tags",
	)

	for i := 0; i < TAG_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO tags(
				id,
				name
			) VALUES(
				?,
				?
			)`,
			id,
			"tag-" + string(id),
		)
	}

	database.Instance.MustExec(
		"TRUNCATE TABLE note_tag",
	)

	for i := 0; i < NOTE_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO note_tag(
				id,
				note_id,
				tag_id
			) VALUES(
				?,
				?,
				?
			)`,
			id,
			id,
			(id % (TAG_COUNT + 1)),
		)
	}
}