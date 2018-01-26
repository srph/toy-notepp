package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/macaron.v1"
	"github.com/rs/cors"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"github.com/srph/toy-notepp/models"
	"github.com/srph/toy-notepp/lib/authee"
	"github.com/srph/toy-notepp/routes/notes"
	"github.com/srph/toy-notepp/routes/user"
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
	m.Use(cors.Default().HandlerFunc)
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "static/static",
		Extensions: []string{".html"},
	}))
	m.Use(session.Sessioner())
	m.Use(authee.Macaron())
	m.Group("/api", func() {
		m.Get("/notes", notes.Index)
		m.Post("/notes", binding.Json(notes.CreateForm{}), notes.Create)
		m.Get("/notes/:id", notes.Show)
		m.Put("/notes/:id", binding.Json(notes.UpdateForm{}), notes.Update)
		m.Delete("/notes/:id", notes.Destroy)
		m.Get("/me", user.Me)
		m.Post("/login", binding.Json(user.LoginForm{}), user.Login)
		m.Get("/logout", user.Logout)
	})
	m.Run()
	return nil
}