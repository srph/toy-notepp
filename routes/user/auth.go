package user

import (
	"gopkg.in/macaron.v1"
	"github.com/srph/toy-notepp/lib/authee"
)

type LoginForm struct {
	Username string `bind:"Required"`
	Password string `bind:"Required"`
}

func Me(ctx *macaron.Context, auth *authee.Auth) {
	ctx.JSON(200, auth.User)
}

func Login(ctx *macaron.Context, form LoginForm, auth *authee.Auth) {
	err := auth.Login(form.Username, form.Password)
	if err != nil {
		ctx.HTML(400, err.Error())
	} else {
		ctx.JSON(200, auth.User)
	}
}

func Logout(ctx *macaron.Context, auth *authee.Auth) {
	auth.Logout()
	ctx.JSON(200, auth.User)
}