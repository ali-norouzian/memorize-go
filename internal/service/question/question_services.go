package question

import (
	"memorize/internal/model"
	"memorize/internal/repository"
	"memorize/pkg/reflection"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type QuestionService struct {
	repository.IRepository[model.Question]
}

func NewQuestionService(questionRepo repository.IRepository[model.Question]) *QuestionService {
	return &QuestionService{IRepository: questionRepo}
}

func (s *QuestionService) ListQuestions(req repository.PaginateRequest) (*repository.PaginatedResult[ListQuestionsResponse], error) {
	entities, err := s.FindAll(req)
	if err != nil {
		return nil, err
	}

	var resp repository.PaginatedResult[ListQuestionsResponse]
	if err := copier.Copy(&resp, entities); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *QuestionService) GetQuestionByID(entityID uint) (*model.Question, error) {
	entity := &model.Question{Model: gorm.Model{ID: entityID}}
	return entity, s.First(entity)
}

func (s *QuestionService) CreateQuestion(req *CreateQuestionRequest) (*CreateQuestionResponse, error) {
	var entity model.Question
	if err := copier.Copy(&entity, req); err != nil {
		return nil, err
	}

	if err := s.Create(&entity); err != nil {
		return nil, err
	}

	return &CreateQuestionResponse{entity.ID}, nil
}

func (s *QuestionService) UpdateQuestion(req *UpdateQuestionRequest) error {
	if err := s.UpdateFields(&model.Question{}, reflection.StructToMap(*req)); err != nil {
		return err
	}

	return nil
}

func (s *QuestionService) DeleteQuestion(entityID uint) error {
	return s.DeleteByID(entityID)
}
