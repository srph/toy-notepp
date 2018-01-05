package models

import(
	"github.com/jinzhu/gorm"
)

var (
	Instance *gorm.DB
)

func Init() {
	var err error
	Instance, err = gorm.Open(
		"mysql",
		"root@tcp(localhost)/failbook?parseTime=true",
	)
	if err != nil {
		panic(err)
	}
	// defer Instance.Close()
}