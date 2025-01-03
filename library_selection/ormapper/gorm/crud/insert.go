package crud

import (
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"
	"gorm.io/gorm"
)

func Insert(db *gorm.DB) {
	u := models.User{
		Name: "Alice",
	}
	if err := db.Create(&u).Error; err != nil {
		log.Fatal(err)
	}

	p := models.Post{
		Title:   "First Post",
		Content: "Hello GORM",
		UserID:  u.ID,
	}
	if err := db.Create(&p).Error; err != nil {
		log.Fatal(err)
	}

	c := models.Comment{
		Content: "Nice post!",
		UserID:  u.ID,
		PostID:  p.ID,
	}
	if err := db.Create(&c).Error; err != nil {
		log.Fatal(err)
	}
}
