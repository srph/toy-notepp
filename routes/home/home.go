package home

import (
	"gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}