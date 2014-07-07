package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	DB             gorm.DB
	RecordNotFound = gorm.RecordNotFound
)

func init() {
	var err error

	DB, err = gorm.Open("postgres", "dbname=blog_dev sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	// defaults
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().Ping()
}
