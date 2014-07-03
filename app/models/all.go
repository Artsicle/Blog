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

func CreateRecord(i interface{}) (bool, error) {
	if err := DB.Save(i).Error; err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func UpdateRecord(i interface{}) (bool, error) {
	if err := DB.Updates(i).Error; err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func DeleteRecord(i interface{}) (bool, error) {
	if err := DB.Delete(i).Error; err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func FindByMap(m map[string]interface{}, i interface{}) (bool, error) {
	if err := DB.Where(m).Find(i).Error; err != nil {
		return false, err
	} else {
		return true, nil
	}
}
