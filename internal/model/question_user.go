package model

import (
	"time"

	"gorm.io/gorm"
)

var NumberOfDayInEachBox = []uint8{1, 1, 3, 7, 14}

type QuestionUser struct {
	gorm.Model
	UserID       uint      `gorm:"primaryKey"`
	QuestionID   uint      `gorm:"primaryKey"`
	CorrectCount uint8     `gorm:"check:0 < correct_count AND correct_count < 6"`
	BoxNumber    uint8     `gorm:"check:0 < correct_count AND correct_count < 6"`
	ReadyTime    time.Time `gorm:"check:created_at < ready_time"`
}
