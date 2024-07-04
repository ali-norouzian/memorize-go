package user

import (
	"memorize/internal/model/authentication"
	"memorize/internal/repository"
	"memorize/internal/service/authentication/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService *user.UserService
	Validator   *validator.Validate
}

func NewUserHandler(userService *user.UserService,
	validator *validator.Validate) *UserHandler {
	return &UserHandler{UserService: userService, Validator: validator}
}

// @Tags Users
// @Param page query int false "Page"
// @Param page_size query int false "Page size"
// @Router /users [get]
func (ctrl *UserHandler) ListUsers(c *gin.Context) {
	filters := []repository.Filter{}
	if username := c.Query("username"); username != "" {
		filters = append(filters, repository.Filter{Column: "username", Value: username})
	}
	if email := c.Query("email"); email != "" {
		filters = append(filters, repository.Filter{Column: "email", Value: email})
	}

	search := make(map[string]string)
	if username := c.Query("username"); username != "" {
		search["username"] = username
	}
	if email := c.Query("email"); email != "" {
		search["email"] = email
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pagination := repository.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	users, err := ctrl.UserService.ListUsers(repository.PaginateRequest{Filters: filters, Search: search, Pagination: pagination})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Tags Users
// @Param id path int true "id"
// @Router /users/{id} [get]
func (ctrl *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := ctrl.UserService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Param user body users.CreateUserRequest true "entity to create"
// @Router /users [post]
func (ctrl *UserHandler) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctrl.UserService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Tags Users
// @Param id path uint true "id"
// @Param user body users.UpdateUserRequest true "entity to update"
// @Router /users/{id} [put]
func (ctrl *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req user.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)

	if err := ctrl.UserService.UpdateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

// @Tags Users
// @Param id path uint true "id"
// @Router /users/{id} [delete]
func (ctrl *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user := authentication.User{Model: gorm.Model{ID: uint(id)}}

	if err := ctrl.UserService.DeleteUser(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
