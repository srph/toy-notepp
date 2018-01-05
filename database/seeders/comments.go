package seeders

import (
	"github.com/srph/failbook/database"
)

func SeedComments() {
	database.Instance.MustExec(
		"TRUNCATE TABLE comments",
	)

	for i := 0; i < POST_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO comments(
				id,
				post_id,
				user_id,
				content
			) VALUES(
				?,
				?,
				?,
				?
			)`,
			id,
			id,
			(i % USER_COUNT) + 1,
			"bla i'm a cool comment",
		)
	}
}