package seeders

import (
	"fmt"
	"github.com/srph/toy-notepp/database"
)

func SeedGroups() {
	database.Instance.MustExec(
		"TRUNCATE TABLE groups",
	)

	database.Instance.MustExec(
		"TRUNCATE TABLE group_user",
	)

	for i := 0; i < GROUP_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO groups(
				id,
				name,
				description
			) VALUES(
				?,
				?,
				?
			)`,
			id,
			fmt.Sprintf("Ex-Battalion Version %d", id),
			"We don't die, we steal beats oh oh",
		)
	}

	for i := 0; i < USER_COUNT * GROUP_COUNT; i++ {
		id := i + 1

		database.Instance.MustExec(`
			INSERT INTO group_user(
				id,
				group_id,
				user_id
			) VALUES(
				?,
				?,
				?
			)`,
			id,
			id % GROUP_COUNT,
			id % USER_COUNT,
		)
	}
}