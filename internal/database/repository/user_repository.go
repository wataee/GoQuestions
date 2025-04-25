package repository

import (
	"gorm.io/gorm"
	"github.com/wataee/GoQuestions/internal/user"
)

type UserRepository interface {
	CreateUser(user user.User) error
 }

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user user.User) error {
	return r.db.Create(&user).Error
}