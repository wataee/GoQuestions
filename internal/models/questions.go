package models

type QuestionDTO struct {
	ID 					 uint 		`json:"id"`
	QuestionText string 	`json:"question_text"`
}

type AddQuestionDTO struct {
	QuestionText  string     `json:"question_text"`
  CorrectAnswer bool     	 `json:"correct_answer"`
}