package migrations

func MigratePosts() {
	Instance.MustExec(`
		CREATE TABLE posts(
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropPosts() {
	Instance.MustExec("DROP TABLE IF EXISTS posts")
}