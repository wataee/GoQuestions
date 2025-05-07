package repository

import (
	"github.com/wataee/GoQuestions/internal/models"
	"gorm.io/gorm"
)

type QuestionsRepository interface {
	GetQuestions() ([]models.Questions, error)
}

type questionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionsRepository {
	return &questionsRepository{db: db}
}

func (r *questionsRepository) GetQuestions() ([]models.Questions, error) {
	var questions []models.Questions

	if err := r.db.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
