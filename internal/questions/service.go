package questions

import (
	"github.com/wataee/GoQuestions/internal/database/repository"
	"github.com/wataee/GoQuestions/internal/models"
)

type QuestionsService interface {
	QuestionsList() ([]models.QuestionDTO, error)
}

type questionsService struct {
	repo repository.QuestionsRepository
}

func NewQuestionsService(repo repository.QuestionsRepository) QuestionsService {
	return &questionsService{repo: repo}
}

func (s *questionsService) QuestionsList() ([]models.QuestionDTO, error) {
	questions, err := s.repo.GetQuestions()
	if err != nil {
		return nil, err
	}

	var questionsDTO []models.QuestionDTO

	for _,q := range questions {
		questionsDTO = append(questionsDTO, models.QuestionDTO{
			ID: q.ID,
			QuestionText: q.QuestionText,
		})
	}
	return questionsDTO, nil
}