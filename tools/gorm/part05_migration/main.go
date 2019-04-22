package main

import (
	"part05_migration/model"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&model.User{},
		&model.Organization{},
	)
}
