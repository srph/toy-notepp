package models

import(
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/srph/toy-notepp/lib/settings"
)

var (
	Instance *gorm.DB
)

func Init() {
	var err error
	Instance, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?parseTime=true",
			settings.DatabaseUser,
			settings.DatabasePassword,
			settings.DatabaseHost,
			settings.DatabaseName,
		),
	)
	if err != nil {
		panic(err)
	}
}