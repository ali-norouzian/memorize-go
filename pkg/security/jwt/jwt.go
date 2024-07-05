package jwt

import (
	"memorize/config"
	"memorize/internal/model/authentication"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte
var expirationTimeInDay int

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJwtService(config *config.Config) {
	jwtKey = []byte(config.Jwt.Secret)
	expirationTimeInDay = config.Jwt.ExpirationTimeInDay
}

func GenerateJwt(user *authentication.User) (string, error) {
	expirationTime := time.Now().AddDate(0, 0, expirationTimeInDay)
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJwt(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
