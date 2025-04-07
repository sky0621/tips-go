package crud

import (
	"fmt"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/models"
	"gorm.io/gorm"
)

func Select(db *gorm.DB) {
	var allUsers []models.User
	if err := db.Find(&allUsers).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の全件取得 --")
	fmt.Println(allUsers)

	var allUsersWithPostsWithComments []models.User
	if err := db.Preload("Posts.Comments").Find(&allUsersWithPostsWithComments).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- リレーションテーブル込みの全件取得 --")
	fmt.Println(allUsersWithPostsWithComments)

	var userAlice models.User
	if err := db.Where(&models.User{Name: "Alice"}).First(&userAlice).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- 単一テーブル(users)の条件付き取得 --")
	fmt.Println(userAlice)
}
