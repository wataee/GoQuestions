package repository

import (
	"github.com/wataee/GoQuestions/internal/database"
	"github.com/wataee/GoQuestions/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.UserInput) error
	FindByUsername(username string) (bool, error)
 }

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.UserInput) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByUsername(username string) (bool, error) {
	var user database.Users
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil // Если не найден пользователь	
		} else {
			return false, result.Error // Другая ошибка 
		}
	}
	return true, nil // Если найден
}