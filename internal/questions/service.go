package questions

import "github.com/wataee/GoQuestions/internal/database/repository"

type QuestionsService interface {

}

type questionsService struct {
	repo repository.QuestionsRepository
}

func NewQuestionsService(repo repository.QuestionsRepository) QuestionsService {
	return &questionsService{repo: repo}
}