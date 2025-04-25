package database

import "time"

type Users struct {
	ID        uint   	`gorm:"primaryKey;not null"`
	Name      string 	`size:"128;not null"`
	Answers   int
	CreatedAt time.Time
}

type Questions struct {
	ID uint 							`gorm:"primaryKey;not null"`
	QuestionText string  	`gorm:"type:text;not null"`
	CreatedAt time.Time
}