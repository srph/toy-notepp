package posts

import (
	"gopkg.in/macaron.v1"
	"github.com/srph/failbook/models"
	"github.com/srph/failbook/lib/authee"
)

type CreateForm struct {
	Content string `binding:"Required"`
}

type UpdateForm struct {
	Content string `binding:"Required"`
}

func Index(ctx *macaron.Context, auth *authee.Auth) {
	posts := []models.Post{}
	
	models.Instance.Model(&auth.User).
		Order("id desc").
		Related(&posts)

	ctx.JSON(200, map[string]interface{}{
		"data": posts,
	})
}

func Create(ctx *macaron.Context, form CreateForm, auth *authee.Auth) {
	post := models.Post{
		User: auth.User,
		Content: form.Content,
	}

	models.Instance.Create(&post)

	ctx.JSON(200, map[string]interface{}{
		"data": post,
	})
}

func Show(ctx *macaron.Context) {
	id := ctx.Params(":id")
	post := models.Post{}
	models.Instance.Where("id = ?", id).First(&post)

	ctx.JSON(200, map[string]interface{}{
		"data": post,
	})
}

func Update(ctx *macaron.Context, form UpdateForm) {
	id := ctx.Params(":id")

	post := models.Post{}
	models.Instance.Where("id = ?", id).First(&post)
	models.Instance.Model(&post).Update("content", form.Content)

	ctx.JSON(200, map[string]interface{}{
		"data": post,
	})
}

func Destroy(ctx *macaron.Context) {
	id := ctx.Params(":id")	

	post := models.Post{}
	models.Instance.Where("id = ?", id).First(&post)
	models.Instance.Delete(&post)

	ctx.JSON(200, map[string]interface{}{})
}