package seeders

import (
	"github.com/srph/failbook/database"
)

func SeedPosts() {
	database.Instance.MustExec(
		"TRUNCATE TABLE posts",
	)

	for i := 0; i < POST_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO posts(
				id,
				group_id,
				user_id,
				content
			) VALUES(
				?,
				?,
				?,
				?
			)`,
			id,
			(i % 6),
			(i % 5) + 1,
			"bla i'm a cool post",
		)
	}
}