package models

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:255;not null" json:"username"`
	Email string `gotm:"size:100;unique;not null;" json:"email"`
}

func AutoMigrateUser(db *gorm.DB) error{
	return db.AutoMigrate(&User{})
}