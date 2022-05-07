package seed

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lzszq/AssessmentRestfulApi/api/models"
)

func Load(db *gorm.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	var users = []models.User{
		models.User{
			Nickname:  os.Getenv("InitNickname"),
			Email:     os.Getenv("InitEmail"),
			Password:  os.Getenv("InitPassword"),
			AvatarUrl: os.Getenv("InitAvatarUrl"),
		},
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			fmt.Printf("cannot seed users table: %v, maybe already exists!\n", err)
		}
	}
}
