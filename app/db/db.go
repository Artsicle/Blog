package db

import "github.com/jinzhu/gorm"
import _ "github.com/lib/pq"
import "fmt"
import "reflect"
import "Blog/app/models"

var DB gorm.DB

func init() {
  var err error

  DB, err := gorm.Open("postgres", "dbname=blog_dev sslmode=disable")
  if err != nil {
    panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
  }

  // defaults
  DB.DB().SetMaxIdleConns(10)
  DB.DB().SetMaxOpenConns(100)
  DB.DB().Ping()

  for _, model := range []interface{}{models.User{}, models.Post{}} {
    if err := DB.AutoMigrate(model).Error; err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
    }
  }
}
