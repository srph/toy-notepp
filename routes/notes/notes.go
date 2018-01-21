package notes

import (
	"gopkg.in/macaron.v1"
	"github.com/srph/toy-notepp/models"
	"github.com/srph/toy-notepp/lib/authee"
)

type CreateForm struct {
	Content string `binding:"Required"`
}

type UpdateForm struct {
	Content string `binding:"Required"`
}

func Index(ctx *macaron.Context, auth *authee.Auth) {
	notes := []models.Note{}
	
	models.Instance.Model(&auth.User).
		Order("updated_at desc").
		Related(&notes)

	ctx.JSON(200, map[string]interface{}{
		"data": notes,
	})
}

func Create(ctx *macaron.Context, form CreateForm, auth *authee.Auth) {
	note := models.Note{
		User: auth.User,
		Content: form.Content,
	}

	models.Instance.Create(&note)

	ctx.JSON(200, map[string]interface{}{
		"data": note,
	})
}

func Show(ctx *macaron.Context) {
	id := ctx.Params(":id")
	
	note := models.Note{}

	models.Instance.Where("id = ?", id).
		Preload("Tags").
		First(&note)

	ctx.JSON(200, map[string]interface{}{
		"data": note,
	})
}

func Update(ctx *macaron.Context, form UpdateForm) {
	id := ctx.Params(":id")

	note := models.Note{}
	models.Instance.Where("id = ?", id).First(&note)
	models.Instance.Model(&note).Update("content", form.Content)

	ctx.JSON(200, map[string]interface{}{
		"data": note,
	})
}

func Destroy(ctx *macaron.Context) {
	id := ctx.Params(":id")	

	note := models.Note{}
	models.Instance.Where("id = ?", id).First(&note)
	models.Instance.Delete(&note)

	ctx.JSON(200, map[string]interface{}{})
}