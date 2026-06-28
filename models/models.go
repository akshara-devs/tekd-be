package models

import (
	"time"
	"gorm.io/gorm"
)

// Define the database models here.
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Name      string         `json:"name"`
}

// Migrate creates or updates database tables for all models in this file.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}
