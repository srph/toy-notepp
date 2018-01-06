package user

import (
	// "log"
	"gopkg.in/macaron.v1"
	"github.com/srph/failbook/utils/auth"
)

type LoginForm struct {
	Username string `bind:"Required"`
	Password string `bind:"Required"`
}

func Me(ctx *macaron.Context) {
	ctx.JSON(200, auth.User)
}

func Login(ctx *macaron.Context, form LoginForm) {
	err := auth.Login(form.Username, form.Password)
	if err != nil {
		ctx.HTML(400, "NAH")
	} else {
		ctx.JSON(200, auth.User)
	}
}

func Logout(ctx *macaron.Context) {
	auth.Logout()
	ctx.JSON(200, auth.User)
}