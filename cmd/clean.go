package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/toy-notepp/database"
	"github.com/srph/toy-notepp/database/migrations"
)

var DbClean = cli.Command{
	Name: "db:clean",
	Usage: "Drop migrations",
	Action: runDbClean,
}

func runDbClean(c *cli.Context) error {
	database.Init()
	migrations.DropUsers()
	migrations.DropNotes()
	migrations.DropTags()
	migrations.DropNoteTag()
	return nil
}