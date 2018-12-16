package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           string `gorm:primary_key`
	CreditCardID string
	CreditCard   CreditCard `gorm:"foreignkey:ID;association_foreignkey:CreditCardID"`
}

type CreditCard struct {
	ID     string `gorm:primary_key`
	Number string
}

type Result struct {
	ID           string
	CreditCardID string
	Number       string
}

func main() {
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	//db.AutoMigrate(&User{}, &CreditCard{})
	//
	//db.Create(
	//	&User{
	//		ID:           "u1",
	//		CreditCardID: "c1",
	//		CreditCard: CreditCard{
	//			ID:     "c1",
	//			Number: "xxxx-xxxx-xxxx-xxxx",
	//		},
	//	},
	//)

	r := &Result{}
	db.Table("users").
		Select("users.id, users.credit_card_id, credit_cards.number").
		Joins("inner join credit_cards on credit_cards.id = users.credit_card_id").
		First(r)
	fmt.Printf("%#v\n", r)
}
