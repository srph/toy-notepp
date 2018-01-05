package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/srph/failbook/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Failbook"
	app.Usage = "Yada yada yada, hello world!"
	app.Version = "1"
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.DbMigration,
		cmd.DbClean,
		cmd.DbSeed,
	}
	app.Run(os.Args)
}