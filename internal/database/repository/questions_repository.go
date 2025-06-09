package repository

import (
	"github.com/wataee/GoQuestions/internal/models"
	"gorm.io/gorm"
	"errors"
)

type QuestionsRepository interface {
	GetQuestions() ([]models.Questions, error)
	AddQuestion(question models.AddQuestionDTO) (error)
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

func (r *questionsRepository)AddQuestion(dto models.AddQuestionDTO) (error) {
	var existing models.Questions
	err := r.db.Where("question_text = ?", dto.QuestionText).First(&existing).Error
	if err == nil {
		return errors.New("вопрос в таблице уже существует")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	

	question := models.Questions{
		QuestionText: dto.QuestionText,
		CorrectAnswer: dto.CorrectAnswer,
	}

	return r.db.Create(&question).Error
}