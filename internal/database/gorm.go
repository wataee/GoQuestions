package database

import (
	"log"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	
	"github.com/wataee/GoQuestions/config"
)

func ConnectDB() (*gorm.DB, error) {
	// cfg := config.GetDBConfig()

	dsn := config.DbURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Не удалось подключиться к БД: %v", err)
		return nil, err
	}

	log.Println("Успешное подключение к БД!")
	return db, nil
}