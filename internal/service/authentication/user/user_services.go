package user

import (
	"memorize/internal/model/authentication"
	"memorize/internal/repository"
	"memorize/pkg/reflection"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService struct {
	repository.IRepository[authentication.User]
}

func NewUserService(userRepo repository.IRepository[authentication.User]) *UserService {
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

func (s *UserService) GetUserByID(userID uint) (*authentication.User, error) {
	user := &authentication.User{Model: gorm.Model{ID: userID}}
	return user, s.First(user)
}

func (s *UserService) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	var user authentication.User
	if err := copier.Copy(&user, req); err != nil {
		return nil, err
	}

	if err := s.Create(&user); err != nil {
		return nil, err
	}

	return &CreateUserResponse{user.ID}, nil
}

func (s *UserService) UpdateUser(req *UpdateUserRequest) error {
	if err := s.UpdateFields(&authentication.User{}, reflection.StructToMap(*req)); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(userID uint) error {
	return s.DeleteByID(userID)
}
