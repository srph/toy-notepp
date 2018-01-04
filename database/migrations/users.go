package migrations

func MigrateUsers() {
	Instance.MustExec(`
		CREATE TABLE users(
			id INT NOT NULL AUTO_INCREMENT,
			username VARCHAR(100) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE (username)
		)
	`)
}

func DropUsers() {
	Instance.MustExec("DROP TABLE IF EXISTS users")
}