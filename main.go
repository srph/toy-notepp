package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/toy-notepp/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Note++"
	app.Usage = "A toy note-taking web app"
	app.Version = "1"
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.DbMigration,
		cmd.DbClean,
		cmd.DbSeed,
	}
	app.Run(os.Args)
}