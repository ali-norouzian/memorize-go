package box

import (
	"memorize/internal/model"
	"memorize/internal/repository"
)

type BoxService struct {
	QuestionUserRepo repository.IRepository[model.QuestionUser]
	QuestionRepo     repository.IRepository[model.Question]
}

func NewBoxService(questionUserRepo repository.IRepository[model.QuestionUser],
	questionRepo repository.IRepository[model.Question]) *BoxService {
	return &BoxService{QuestionUserRepo: questionUserRepo,
		QuestionRepo: questionRepo}
}

func (s *BoxService) GetUserBoxesStatus(userId uint) (*GetUserBoxesStatusResponse, error) {
	whereCondition := model.QuestionUser{UserID: userId}
	var questionUsers []model.QuestionUser
	err := s.QuestionUserRepo.FindAlls(&questionUsers, &whereCondition)
	if err != nil {
		return nil, err
	}

	if len(questionUsers) == 0 {
		questionCount, err := s.QuestionRepo.Count(&model.Question{})
		if err != nil {
			return nil, err
		}

		return &GetUserBoxesStatusResponse{
			ActiveCardsCount:    []uint{uint(*questionCount), 0, 0, 0, 0, 0},
			InWaitingCardsCount: []uint{0, 0, 0, 0, 0, 0},
		}, nil

	}

	// entities, err := s.FindAll(req)
	// if err != nil {
	// 	return nil, err
	// }

	// var resp repository.PaginatedResult[ListQuestionsResponse]
	// if err := copier.Copy(&resp, entities); err != nil {
	// 	return nil, err
	// }

	return &GetUserBoxesStatusResponse{}, nil
}
