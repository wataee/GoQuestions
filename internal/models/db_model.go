package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`
	Username  string    `gorm:"size:128;not null"`
	Password  string    `gorm:"size:256;not null"`
	Role      string    `gorm:"size:64;default:'user'"`
	Answers   int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

type Question struct {
    ID            uint       `gorm:"primaryKey;not null"`
    QuestionText  string     `gorm:"type:text;not null"`
    CorrectAnswer bool     	 `gorm:"not null"`
    CreatedAt     time.Time  `gorm:"autoCreateTime;not null"`
    UpdatedAt     time.Time  `gorm:"autoUpdateTime;not null"`
}
