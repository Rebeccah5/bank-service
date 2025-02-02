package config

import (
	"log"
	"os"

	"bank-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the DATABASE_URL from the environment variables
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("database url not set")
	}

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run Migrations
	err = db.AutoMigrate(&models.Account{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	DB = db
}

func InitalBalance() {
	var count int64
	DB.Model(&models.Account{}).Count(&count)
	if count == 0 {
		DB.Create(&models.Account{Balance: 100000}) // Initial balance
		log.Println("initial account balance")
	}
}
