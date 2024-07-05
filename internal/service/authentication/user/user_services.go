package user

import (
	"memorize/internal/model"
	"memorize/internal/repository"
	"memorize/pkg/reflection"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService struct {
	repository.IRepository[model.User]
}

func NewUserService(userRepo repository.IRepository[model.User]) *UserService {
	return &UserService{IRepository: userRepo}
}

func (s *UserService) ListUsers(req repository.PaginateRequest) (*repository.PaginatedResult[ListUsersResponse], error) {
	entities, err := s.FindAll(req)
	if err != nil {
		return nil, err
	}

	var resp repository.PaginatedResult[ListUsersResponse]
	if err := copier.Copy(&resp, entities); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *UserService) GetUserByID(userID uint) (*model.User, error) {
	user := &model.User{Model: gorm.Model{ID: userID}}
	return user, s.First(user)
}

func (s *UserService) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	var user model.User
	if err := copier.Copy(&user, req); err != nil {
		return nil, err
	}

	if err := s.Create(&user); err != nil {
		return nil, err
	}

	return &CreateUserResponse{user.ID}, nil
}

func (s *UserService) UpdateUser(req *UpdateUserRequest) error {
	if err := s.UpdateFields(&model.User{}, reflection.StructToMap(*req)); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(userID uint) error {
	return s.DeleteByID(userID)
}
