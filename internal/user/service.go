package user

import (
	"github.com/wataee/GoQuestions/internal/models"
	"github.com/wataee/GoQuestions/internal/database/repository"
)

type UserService interface {
	Login(input models.UserInput) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(input models.UserInput) (string, error) {
	return "", nil
}
