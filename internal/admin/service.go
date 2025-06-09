package admin

import (
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/models"
)

// admin.GET("/user_list", )
// admin.DELETE("/delete_user/:id", )
// admin.POST("/addquestion", questionsHandler.QuestionAddHandler)
type AdminService interface {
	UserList() ([]models.ProfileDTO, error)
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
