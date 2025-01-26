package config

import (
	"fmt"
	"log"
	"back-eco-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)


var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	
	var err error 
	DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = models.AutoMigrateUser(DB)

	if err != nil {
		log.Fatalf("Error migrating user model: %v",err)
	}
}