package posts

import (
	"gopkg.in/macaron.v1"
	"github.com/srph/failbook/models"
	"github.com/srph/failbook/utils"
)

type CreateForm struct {
	Content string `binding:"Required"`
}

type UpdateForm struct {
	Content string `binding:"Required"`
}

func Index(ctx *macaron.Context) {
	posts := []models.Post{}
	models.Instance.Order("id desc").Find(&posts)

	ctx.JSON(200, map[string]interface{}{
		"data": posts,
	})
}

func Create(ctx *macaron.Context, form CreateForm) {
	post := models.Post{
		User: utils.AuthenticatedUser(),
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