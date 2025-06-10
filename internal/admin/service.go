package admin

import (
	"strconv"

	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/models"
	
)

type AdminService interface {
	UserList() ([]models.ProfileDTO, error)
	QuestionsAdd(question models.AddQuestionDTO) (error)
	DeleteUser(userID string) (error)
}

type adminService struct {
	userRepo repository.UserRepository
	questionsRepo repository.QuestionsRepository
}

func NewAdminService(userRepo repository.UserRepository, questionsRepo repository.QuestionsRepository) AdminService {
	return &adminService{userRepo: userRepo, questionsRepo: questionsRepo}
}

func (a *adminService) UserList() ([]models.ProfileDTO, error) {
	var usersDTO []models.ProfileDTO
	users, err := a.userRepo.GetUserList()
	if err != nil {
		return usersDTO, err
	}
	for _,u := range users {
		usersDTO = append(usersDTO, models.ProfileDTO{
			ID: u.ID,
			Username: u.Username,
			Role: u.Role,
			Answers: u.Answers,
			CreatedAt: u.CreatedAt,
		})
	}

	return usersDTO, nil
}

func (a *adminService) QuestionsAdd(question models.AddQuestionDTO) (error) {
	err := a.questionsRepo.AddQuestion(question)
	if err != nil {
		return err
	}

	return nil
}

func (a *adminService) DeleteUser(userID string) (error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}
	if err := a.userRepo.DeleteUser(id); err != nil {
		return err
	}

	return nil

}