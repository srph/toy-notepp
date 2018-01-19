package migrations

import (
	"github.com/srph/toy-notepp/database"
)

func MigratePosts() {
	database.Instance.MustExec(`
		CREATE TABLE posts(
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL,
			group_id INT,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		)
	`)
}

func DropPosts() {
	database.Instance.MustExec("DROP TABLE IF EXISTS posts")
}