package config

import (
	"fmt"
	"go-repaso/internal/domain"
	"go-repaso/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&domain.Task{},
	)
	if err != nil {
		log.Fatal("Error migrando la BD:", err)
	}

	DB = db
	fmt.Println("Base de datos conectada")
	return DB
}
