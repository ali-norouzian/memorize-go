package jwt

import (
	"memorize/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Jwt struct {
	JwtKey              []byte
	expirationTimeInDay int
}

func NewJwt(config *config.Config) *Jwt {
	return &Jwt{
		JwtKey:              []byte(config.Jwt.Secret),
		expirationTimeInDay: config.Jwt.ExpirationTimeInDay,
	}
}

func (jwtInstance *Jwt) GenerateJwt(claims *Claims) (string, error) {
	// add expire time here for them
	expirationTime := time.Now().AddDate(0, 0, jwtInstance.expirationTimeInDay)
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtInstance.JwtKey)
}

func (jwtInstance *Jwt) VerifyJwt(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtInstance.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
