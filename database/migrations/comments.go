package migrations

func MigrateComments() {
	Instance.MustExec(`
		CREATE TABLE comments(
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL,
			post_id INT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropComments() {
	Instance.MustExec("DROP TABLE IF EXISTS comments")
}