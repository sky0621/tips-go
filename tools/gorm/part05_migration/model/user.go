package model

type User struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Birthday int
}
