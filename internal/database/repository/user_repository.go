package repository

import (
	"gorm.io/gorm"
	
	"github.com/wataee/GoQuestions/internal/models"
)

type UserRepository interface {
	CreateUser(user models.UserInput) (int,error)
	GetByUsername(username string) (models.Users, error)
	GetByID(userID int) (models.Users, error)
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
		return 0, err
	}
	return int(user.ID), nil
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

func (r *userRepository) GetByID(userID int) (models.Users, error) {
	var user models.Users
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return models.Users{}, err
	}
	return user, nil
}