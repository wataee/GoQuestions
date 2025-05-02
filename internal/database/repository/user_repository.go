package repository

import (
	"fmt"

	"github.com/wataee/GoQuestions/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.UserInput) (int,error)
	FindByUsername(username string) (bool, error)
	GetByUsername(username string) (models.Users, error)
 }

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(input models.UserInput) (int, error) {
	user := models.Users{
    Username: input.Username,
    Password: input.Password,
    Role:     input.Role,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании юзера в БД: %v", err)
	}
	return int(user.ID), nil
}


func (r *userRepository) FindByUsername(username string) (bool, error) {
	var user models.Users
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

func (r *userRepository) GetByUsername(username string) (models.Users, error) {
	var user models.Users
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Users{}, gorm.ErrRecordNotFound
		}
		return models.Users{}, err
	}
	return user, nil
}