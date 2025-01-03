package migration

import (
	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
