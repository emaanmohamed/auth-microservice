package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/auth-microservice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

func Migrate() {
	db.AutoMigrate(&Token{})
}
