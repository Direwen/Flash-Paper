package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/direwen/flashpaper/internal/models"
)

// Define & Export Global DB instance
// Use Pointer to share memory and prevent duplicating when called every time
var DB *gorm.DB

func ConnectDB() {
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	//Connect to DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	//Migrate models
	log.Println("Running Migrations")
	err = DB.AutoMigrate(&models.User{}, &models.Snippet{})
	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}
	log.Println("Migrations ran successfully")

	log.Println("Connected to database successfully")
}

func GetDB() *gorm.DB {
	return DB
}
