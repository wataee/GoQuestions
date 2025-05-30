package database

import (
	"log"
	"fmt"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	
	"github.com/wataee/GoQuestions/config"
)

func ConnectDB() (*gorm.DB, error) {
	cfg := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Не удалось подключиться к БД: %v", err)
		return nil, err
	}

	log.Println("Успешное подключение к БД!")
	return db, nil
}