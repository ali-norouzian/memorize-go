package question

type CreateQuestionRequest struct {
	Title        *string `json:"title" validate:"omitempty,max=50"`
	QuestionText string  `json:"question_text" validate:"required,max=100"`
	AnswerText   string  `json:"answer_text" validate:"required,max=1000"`
}

type UpdateQuestionRequest struct {
	ID           uint    `json:"-"`
	Title        *string `json:"title" validate:"omitempty,max=50"`
	QuestionText string  `json:"question_text" validate:"required,max=100"`
	AnswerText   string  `json:"answer_text" validate:"required,max=1000"`
}
