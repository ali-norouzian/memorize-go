package authentication

import (
	"memorize/internal/service/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	Validator   *validator.Validate
	AuthService *authentication.AuthService
}

func NewAuthHandler(validator *validator.Validate,
	authService *authentication.AuthService) *AuthHandler {
	return &AuthHandler{Validator: validator, AuthService: authService}
}

// @Tags Authentication
// @Param user body authentication.RegisterUserRequest true "register"
// @Router /auth/register [post]
func (hndlr *AuthHandler) RegisterUser(c *gin.Context) {
	var req authentication.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hndlr.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := hndlr.AuthService.RegisterUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Tags Authentication
// @Param user body authentication.LoginUserRequest true "login"
// @Router /auth/login [post]
func (hndlr *AuthHandler) LoginUser(c *gin.Context) {
	var req authentication.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hndlr.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := hndlr.AuthService.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
