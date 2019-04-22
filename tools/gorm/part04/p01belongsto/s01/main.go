package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string  `gorm:"primary_key"`
	Profile   Profile `gorm:"foreignkey:ProfileID;association_foreignkey:Refer"`
	ProfileID string
}

type Profile struct {
	ID    string `gorm:"primary_key"`
	Refer string
	Name  string
}

type Result struct {
	ID        string
	ProfileID string
	Name      string
}

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User{}, &Profile{})

	db.Create(
		&User{
			ID: "1",
			Profile: Profile{
				ID:    "1",
				Refer: "1",
				Name:  "dummy",
			},
			ProfileID: "1",
		},
	)

	r := &Result{}
	db.Table("users").
		Select("users.id, users.profile_id, profiles.name").
		Joins("inner join profiles on profiles.id = users.profile_id").
		First(r)
	fmt.Printf("%#v\n", r)
}
