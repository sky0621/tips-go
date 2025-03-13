package main

import (
	"fmt"
	"reflect"
)

type user struct {
	ID   int64
	Name string
}

// 新しい型を定義する（元の型から独立した型になる）
type email string

// 既存の型 user のエイリアスを作る
type user2 = user

func main() {
	strEmail := "dummy@example.com"
	email := email(strEmail)

	t1 := reflect.TypeOf(strEmail)
	t2 := reflect.TypeOf(email)
	fmt.Println(t1 == t2)

	u1 := user{ID: 1, Name: "user1"}
	u2 := user2{ID: 2, Name: "user2"}
	tu1 := reflect.TypeOf(u1)
	tu2 := reflect.TypeOf(u2)
	fmt.Println(tu1 == tu2)
}
