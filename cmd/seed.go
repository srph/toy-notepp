package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/toy-notepp/database"
	"github.com/srph/toy-notepp/database/seeders"
)

var DbSeed = cli.Command{
	Name: "db:seed",
	Usage: "Run seeders",
	Action: runSeeders,
}

func runSeeders(c *cli.Context) error {
	database.Init()
	seeders.SeedUsers()
	seeders.SeedPosts()
	seeders.SeedComments()
	seeders.SeedGroups()
	return nil
}