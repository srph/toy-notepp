package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/macaron.v1"
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
	m := macaron.Classic()
	m.Get("/", home)
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	m.Run()
	return nil
}

func home(ctx *macaron.Context) {
	response := map[string]interface{}{
		"pogi": "kier",
		"age": 69,
	}

	ctx.JSON(200, response)
}