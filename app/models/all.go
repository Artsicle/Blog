package models

import (
	. "blog/app/db"
	"fmt"
	"github.com/revel/revel"
	"reflect"
)

func init() {
	revel.OnAppStart(AutoMigrate)
}

func AutoMigrate() {
	for _, model := range []interface{}{User{}, Post{}} {
		if err := DB.AutoMigrate(model).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
		}
	}
}
