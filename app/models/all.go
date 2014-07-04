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

func CreateRecord(i interface{}) error {
	if err := DB.Save(i).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRecord(i interface{}) error {
	if err := DB.Updates(i).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRecord(i interface{}) error {
	if err := DB.Delete(i).Error; err != nil {
		return err
	}
	return nil
}

// boolean nf is true if we want to report RecordNotFound errors
func FindByMap(m map[string]interface{}, i interface{}, nf bool) error {
	err := DB.Where(m).Find(i).Error
	if err != nil && (nf || err != RecordNotFound) {
		return err
	}
	return nil
}

// boolean nf is true if we want to report RecordNotFound errors
func FindAll(i interface{}, nf bool) error {
	err := DB.Find(i).Error
	if err != nil && (nf || err != RecordNotFound) {
		return err
	}
	return nil
}
