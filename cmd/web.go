package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/macaron.v1"
	"github.com/srph/failbook/models"
)

var Web = cli.Command{
	Name: "web",
	Usage: "start web server",
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "port, p",
			Value: "3000",
			Usage: "Assign a custom port number",
		},
	},
}

func runWeb(c *cli.Context) error {
	models.Init()
	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	m.Get("/", home)
	m.Run()
	return nil
}

func home(ctx *macaron.Context) {
	users := []models.User{}
	models.Instance.Find(&users)

	response := map[string]interface{}{
		"data": users,
	}

	ctx.JSON(200, response)
}