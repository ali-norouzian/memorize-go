package authentication

import (
	"memorize/internal/model/authentication"
	"memorize/internal/repository"
	"strings"

	"github.com/jinzhu/copier"
)

type AuthService struct {
	repository.IRepository[authentication.User]
}

func NewAuthService(userRepo repository.IRepository[authentication.User]) *AuthService {
	return &AuthService{IRepository: userRepo}
}

func (srvc *AuthService) RegisterUser(req *RegisterUserRequest) (*RegisterUserResponse, error) {
	var user authentication.User
	if err := copier.Copy(&user, req); err != nil {
		return nil, err
	}
	user.Username = strings.Split(user.Email, "@")[0]

	if err := srvc.Create(&user); err != nil {
		return nil, err
	}

	var resp RegisterUserResponse
	if err := copier.Copy(&resp, user); err != nil {
		return nil, err
	}

	return &resp, nil
}
