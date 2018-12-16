package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string  `gorm:primary_key`
	Profile   Profile `gorm:"foreignkey:ProfileID;association_foreignkey:Refer"`
	ProfileID string
}

type Profile struct {
	ID    string `gorm:primary_key`
	Refer string
	Name  string
}

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	//db.AutoMigrate(&User{}, &Profile{})
	//
	//db.Create(
	//	&User{
	//		ID: "1",
	//		Profile: Profile{
	//			ID:    "1",
	//			Refer: "1",
	//			Name:  "dummy",
	//		},
	//		ProfileID: "1",
	//	},
	//)
	//
	//profile := Profile{}
	//profile.ID = "1"
	//user := User{Profile: profile}
	//user.ID = "1"
	//user.ProfileID = "1"
	//
	//db.Model(&user).Related(&profile)
	//fmt.Printf("user: {ID: %v, ProfileID: %v, Profile: {ID: %v, Name: %v}}\n", user.ID, user.ProfileID, user.Profile.ID, user.Profile.Name)
	//users := []*User{}
	//db.Find(&users)
	//for _, user := range users {
	//	fmt.Printf("user: {ID: %v, Prifile: {ID: %v, Name: %v}}\n", user.ID, user.Profile.ID, user.Profile.Name)
	//}
	rows, err := db.Table("users").Select("users.id, users.profile_id, profiles.name").Joins("inner join profiles on profiles.id = users.profile_id").Rows()
	if err != nil {
		panic(err)
	}
	//bl := rows.NextResultSet()
	//fmt.Println(bl)
	for rows.Next() {
		fmt.Println("a")
		var re *Result
		rows.Scan(&re)
		fmt.Println(re)
	}
	var results []*Result
	rows.Scan(&results)
	for _, res := range results {
		fmt.Printf("user: {ID: %v, Prifile: {ID: %v, Name: %v}}\n", res.ID, res.ProfileID, res.Name)
	}

	u1 := &User{}
	db.Find(u1)
	fmt.Printf("%#v\n", u1)

	p1 := &Profile{}
	db.Find(p1)
	fmt.Printf("%#v\n", p1)

}

type Result struct {
	ID        string
	ProfileID string
	Name      string
}
