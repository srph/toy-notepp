package notes

import (
	"gopkg.in/macaron.v1"
	"strings"
	"github.com/srph/toy-notepp/models"
	"github.com/srph/toy-notepp/lib/authee"
	// "github.com/srph/toy-notepp/lib/utils"
)

type CreateForm struct {
	Content string `binding:"Required"`
}

type UpdateForm struct {
	Content string `binding:"Required"`
	Tags []string
}

type DeleteForm struct {
	Force bool
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

	models.Instance.Where("id = ?", id).
		Preload("Tags").
		First(&note)

	models.Instance.Save(&note)

	// Blow away all existing tags
	models.Instance.Model(&note).
		Association("Tags").
		Clear()

	// Create all tags
	var tags []models.Tag
	for _, name := range form.Tags {
		tag := models.Tag{}
		models.Instance.FirstOrCreate(&tag, models.Tag{Name: strings.Trim(name, " ") })
		tags = append(tags, tag)
	}

	// Relate all tags
	models.Instance.Model(&note).
		Association("Tags").
		Append(tags)

	ctx.JSON(200, map[string]interface{}{
		"data": note,
	})
}

func Destroy(ctx *macaron.Context, form DeleteForm) {
	id := ctx.Params(":id")

	if (form.Force) {
		models.Instance.Where("id = ?", id).Delete(models.Note{})
	} else {
		models.Instance.Unscoped().Where("id = ?", id).Delete(models.Note{})
	}

	ctx.JSON(200, map[string]interface{}{})
}