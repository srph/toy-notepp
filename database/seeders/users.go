package seeders

import (
	"fmt"
	"github.com/srph/failbook/database"
	"github.com/srph/failbook/lib/utils"
)

func SeedUsers() {
	password := utils.Bcrypt("123")

	database.Instance.MustExec(
		"TRUNCATE TABLE users",
	)

	for i := 0; i < 5; i++ {
		database.Instance.MustExec(`
			INSERT INTO users(
				id,
				username,
				password
			) VALUES(
				?,
				?,
				?
			)`,
			i + 1,
			fmt.Sprintf("beyan%d", i + 1),
			password,
		)
	}
}