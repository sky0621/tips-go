package main

import (
	"tips-go/tools/gorm/part02/model"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(
		&model.HogeFuga{},
	)
}
