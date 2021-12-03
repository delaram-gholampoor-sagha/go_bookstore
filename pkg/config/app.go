package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// To import a package solely for its side-effects (initialization), use the blank identifier as explicit package name:

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
