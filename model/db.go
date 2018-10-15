package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/makki0205/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	fmt.Println(config.Env("DB_USER"))
	db, err := gorm.Open("mysql", config.Env("DB_USER")+":"+config.Env("DB_PASS")+"@tcp("+config.Env("DB_CONN")+")/kanban?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todo{})

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

