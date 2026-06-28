package pkg

import (
	"log"

	"github.com/akshara-devs/tekd-be/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := models.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	log.Println("[info] database connected and migrated")
}
