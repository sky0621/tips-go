package crud

import (
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) {
	if err := db.Where("id = ?", 1).Delete(&models.User{}).Error; err != nil {
		log.Fatal(err)
	}
}
