package authentication

import (
	"errors"
	"memorize/internal/model/authentication"
	"memorize/internal/repository"
	"memorize/pkg/security/hash"
	"memorize/pkg/security/jwt"
	"strings"

	"github.com/jinzhu/copier"
)

type AuthService struct {
	repository.IRepository[authentication.User]
	*jwt.Jwt
}

func NewAuthService(userRepo repository.IRepository[authentication.User],
	jwt *jwt.Jwt) *AuthService {
	return &AuthService{IRepository: userRepo, Jwt: jwt}
}

func (srvc *AuthService) RegisterUser(req *RegisterUserRequest) (*RegisterUserResponse, error) {
	var user authentication.User
	if err := copier.Copy(&user, req); err != nil {
		return nil, err
	}
	user.Username = strings.Split(user.Email, "@")[0]
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := srvc.Create(&user); err != nil {
		return nil, err
	}

	var resp RegisterUserResponse
	if err := copier.Copy(&resp, user); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (srvc *AuthService) LoginUser(req *LoginUserRequest) (*LoginUserResponse, error) {
	user := authentication.User{
		Username: req.Username,
	}
	if err := srvc.First(&user); err != nil {
		return nil, err
	}

	if !hash.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	claims := jwt.Claims{
		UserID:   user.ID,
		Username: user.Username,
	}
	token, err := srvc.GenerateJwt(&claims)
	if err != nil {
		return nil, err
	}

	resp := LoginUserResponse{
		Username: user.Username,
		Token:    token,
	}

	return &resp, nil
}
