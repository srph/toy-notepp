package migrations

func MigrateGroupUser() {
	Instance.MustExec(`
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
	Instance.MustExec("DROP TABLE IF EXISTS users")
}