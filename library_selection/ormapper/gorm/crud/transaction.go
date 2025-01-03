package crud

import (
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"

	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		var p models.Post
		if err := tx.First(&p).Error; err != nil {
			return err
		}
		if err := tx.Model(&p).Update("Title", "Second Post").Error; err != nil {
			return err
		}
		//if p.ID == 1 {
		//	return errors.New("ERR")
		//}
		if err := tx.Model(&p).Update("Title", "Third Post").Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println("transaction err:", err)
	}
}
