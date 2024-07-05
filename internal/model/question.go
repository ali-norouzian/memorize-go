package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Title        *string `gorm:"size:50"`
	QuestionText string  `gorm:"size:100;not null"`
	AnswerText   string  `gorm:"size:1000;not null"`
	User         []User  `gorm:"many2many:question_users;"`
}
