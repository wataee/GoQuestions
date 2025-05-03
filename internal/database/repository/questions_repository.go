package repository

import (
	"gorm.io/gorm"

)

type QuestionsRepository interface {

}

type questionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) QuestionRepository {
	return &questionsRepository{db: db}
}