package middleware

import (
	"memorize/pkg/security"
	"memorize/pkg/security/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	*jwt.Jwt
}

func NewAuthMiddleware(jwt *jwt.Jwt) *AuthMiddleware {
	return &AuthMiddleware{Jwt: jwt}
}

func (authMiddleware *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims, err := authMiddleware.VerifyJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set(security.UsernameContextKey, claims.Username)
		c.Set(security.UserIdContextKey, claims.UserID)

		c.Next()
	}
}
