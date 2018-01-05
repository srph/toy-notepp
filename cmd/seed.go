package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/failbook/database/seeders"
)

var DbSeed = cli.Command{
	Name: "db:seed",
	Usage: "Run seeders",
	Action: runSeeders,
}

func runSeeders(c *cli.Context) error {
	seeders.SeedUsers()
	seeders.SeedPosts()
	seeders.SeedComments()
	seeders.SeedGroups()
	return nil
}