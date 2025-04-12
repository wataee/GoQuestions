package database

import (
	"fmt"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	
	"goquestions/config"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключится к бд" + err.Error())
	}

	DB = db
	fmt.Println("Успешное подключение к БД!")
}