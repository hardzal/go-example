package models

import (
	"github.com/hardzal/go-example/go-books/pkg/config"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&User{})
}
