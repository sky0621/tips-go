package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Posts []Post
}

type Post struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Title     string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Comments []Comment
}

type Comment struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"index"`
	PostID    uint   `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
