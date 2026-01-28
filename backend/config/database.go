package config

import (
	"fmt"
	"log"
	"os"

	"jogokariyan-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_Host"),
		os.Getenv("DB_User"),
		os.Getenv("DB_Password"),
		os.Getenv("DB_Name"),
		os.Getenv("DB_Port"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Database connected successfully")

	// Optional: AutoMigrate to ensure schema is in sync
	DB.AutoMigrate(&models.Peserta{}, &models.Kegiatan{}, &models.Presensi{}, &models.AppSettings{}, &models.AppImage{}, &models.BackgroundImage{})
}
