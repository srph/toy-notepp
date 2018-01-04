package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/failbook/database/migrations"
)

var DbMigration = cli.Command{
	Name: "db:migrate",
	Usage: "Run migrations",
	Action: runMigrations,
}

func runMigrations(c *cli.Context) error {
	migrations.MigrateUsers()
	migrations.MigratePosts()
	return nil
}