package database

import (
	"log"
	"fmt"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	
	"github.com/wataee/GoQuestions/config"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	DB = db
	log.Println("Успешное подключение к БД!")
}