package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"github.com/srph/toy-notepp/models"
	"github.com/srph/toy-notepp/lib/authee"
	"github.com/srph/toy-notepp/routes/posts"
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
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(authee.Macaron())
	m.Get("/", home)
	m.Get("/posts", posts.Index)
	m.Post("/posts", binding.Bind(posts.CreateForm{}), posts.Create)
	m.Get("/posts/:id", posts.Show)
	m.Put("/posts/:id", binding.Bind(posts.UpdateForm{}), posts.Update)
	m.Delete("/posts/:id", posts.Destroy)
	m.Get("/me", user.Me)
	m.Get("/login", binding.Bind(user.LoginForm{}), user.Login)
	m.Get("/logout", user.Logout)
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