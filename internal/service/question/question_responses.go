package question

import "time"

type CreateQuestionResponse struct {
	ID uint
}

type ListQuestionsResponse struct {
	ID           uint
	Title        *string
	QuestionText string
	AnswerText   string
}

type GetQuestionResponse struct {
	ID           uint
	Title        *string
	QuestionText string
	AnswerText   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
