package model

type User struct {
	ID         string `gorm:"primary_key"`
	CreditCard CreditCard
}

type CreditCard struct {
	ID     string `gorm:"primary_key"`
	UserID string
}
