package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Organization struct {
	ID    string `gorm:primary_key`
	Name  string
	Users []User `gorm:"foreignkey:OrganizationID"`
}

type User struct {
	ID             string `gorm:primary_key`
	Name           string
	OrganizationID string
}

type ResOrgUser struct {
	OrganizationID   string
	OrganizationName string
	UserID           string
	UserName         string
}

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	//db.AutoMigrate(&User{}, &Organization{})
	//
	//ts := db.Begin()
	//defer ts.Commit()
	//
	//ts.Create(
	//	&Organization{
	//		ID:   "o01",
	//		Name: "SosikiA",
	//		Users: []User{
	//			{ID: "u01", Name: "Sato", OrganizationID: "o01"},
	//			{ID: "u02", Name: "Takahashi", OrganizationID: "o01"},
	//		},
	//	},
	//)
	//
	//ts.Create(
	//	&Organization{
	//		ID:   "o02",
	//		Name: "SosikiB",
	//		Users: []User{
	//			{ID: "u03", Name: "Kato", OrganizationID: "o02"},
	//			{ID: "u04", Name: "Ishida", OrganizationID: "o02"},
	//			{ID: "u05", Name: "Niwa", OrganizationID: "o02"},
	//		},
	//	},
	//)
	//
	//if err := ts.Error; err != nil {
	//	ts.Rollback()
	//}

	var r []ResOrgUser
	db.Table("organizations").
		Select(
			"organizations.id AS organization_id, " +
				"organizations.name AS organization_name, " +
				"users.id AS user_id, " +
				"users.name AS user_name").
		Joins("inner join users on users.organization_id = organizations.id").
		Find(&r)
	for idx, res := range r {
		fmt.Printf("[%d]%#v\n", idx, res)
	}

}
