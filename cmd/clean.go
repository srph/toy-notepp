package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/failbook/database/migrations"
)

var DbClean = cli.Command{
	Name: "db:clean",
	Usage: "Run migrations",
	Action: runDbClean,
}

func runDbClean(c *cli.Context) error {
	migrations.DropUsers()
	migrations.DropPosts()
	migrations.DropComments()
	migrations.DropGroups()
	migrations.DropGroupUser()
	return nil
}