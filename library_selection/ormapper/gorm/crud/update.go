package crud

import (
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"
	"gorm.io/gorm"
)

func Update(db *gorm.DB) {
	if err := db.Model(&models.User{}).Where("id = ?", 1).Update("Name", "Alice Updated").Error; err != nil {
		log.Fatal(err)
	}
}
