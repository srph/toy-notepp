package main

import (
	"github.com/srph/toy-notepp/lib/settings"
	
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/toy-notepp/cmd"
	"os"
)

func main() {
	settings.Init()

	app := cli.NewApp()
	app.Name = settings.AppName
	app.Usage = settings.AppDescription
	app.Version = settings.AppVersion
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.DbMigration,
		cmd.DbClean,
		cmd.DbSeed,
	}
	app.Run(os.Args)
}