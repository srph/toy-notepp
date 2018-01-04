package migrations

func MigrateGroups() {
	Instance.MustExec(`
		CREATE TABLE groups(
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropGroups() {
	Instance.MustExec("DROP TABLE IF EXISTS groups")
}