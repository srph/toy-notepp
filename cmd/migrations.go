package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/toy-notepp/database"
	"github.com/srph/toy-notepp/database/migrations"
)

var DbMigration = cli.Command{
	Name: "db:migrate",
	Usage: "Run migrations",
	Action: runMigrations,
}

func runMigrations(c *cli.Context) error {
	database.Init()
	migrations.MigrateUsers()
	migrations.MigrateNotes()
	migrations.MigrateTags()
	migrations.MigrateNoteTag()
	return nil
}