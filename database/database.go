package database

import (
	"golang_test_day_07/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=dpg-cse56du8ii6s738t0np0-a.singapore-postgres.render.com user=bookdb_sdlq_user password=0HnkSUU2NZHsEENi8aeT2xEhUlhU13VL dbname=bookdb_sdlq port=5432 TimeZone=Asia/Shanghai"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return
	}

	log.Println("Database connected and migration completed successfully!")
}
