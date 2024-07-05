package question

import (
	"memorize/internal/model"
	"memorize/internal/repository"
	"memorize/internal/service/question"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	QuestionService *question.QuestionService
	Validator       *validator.Validate
}

func NewQuestionHandler(questionService *question.QuestionService,
	validator *validator.Validate) *QuestionHandler {
	return &QuestionHandler{QuestionService: questionService,
		Validator: validator,
	}
}

func (hndlr *QuestionHandler) SetRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	QuestionRoutes := routerGroup.Group("/question")
	{
		QuestionRoutes.GET("", hndlr.ListQuestion)
		QuestionRoutes.GET("/:id", hndlr.GetQuestionByID)
		QuestionRoutes.POST("", hndlr.CreateQuestion)
		QuestionRoutes.PUT("/:id", hndlr.UpdateQuestion)
		QuestionRoutes.DELETE("/:id", hndlr.DeleteQuestion)
	}

	return QuestionRoutes
}

// @Tags Question
// @Param page query int false "Page"
// @Param page_size query int false "Page size"
// @Router /admin/question [get]
func (hndlr *QuestionHandler) ListQuestion(c *gin.Context) {
	filters := []repository.Filter{}
	// if Questionname := c.Query("Questionname"); Questionname != "" {
	// 	filters = append(filters, repository.Filter{Column: "Questionname", Value: Questionname})
	// }
	// if email := c.Query("email"); email != "" {
	// 	filters = append(filters, repository.Filter{Column: "email", Value: email})
	// }

	search := make(map[string]string)
	// if Questionname := c.Query("Questionname"); Questionname != "" {
	// 	search["Questionname"] = Questionname
	// }
	// if email := c.Query("email"); email != "" {
	// 	search["email"] = email
	// }

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pagination := repository.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	Question, err := hndlr.QuestionService.ListQuestions(repository.PaginateRequest{Filters: filters, Search: search, Pagination: pagination})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Question)
}

// @Tags Question
// @Param id path int true "id"
// @Router /admin/question/{id} [get]
func (hndlr *QuestionHandler) GetQuestionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	entity, err := hndlr.QuestionService.GetQuestionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	c.JSON(http.StatusOK, entity)
}

// @Tags Question
// @Param Question body question.CreateQuestionRequest true "entity to create"
// @Router /admin/question [post]
func (hndlr *QuestionHandler) CreateQuestion(c *gin.Context) {
	var req question.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hndlr.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := hndlr.QuestionService.CreateQuestion(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Tags Question
// @Param id path uint true "id"
// @Param Question body question.UpdateQuestionRequest true "entity to update"
// @Router /admin/question/{id} [put]
func (hndlr *QuestionHandler) UpdateQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req question.UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hndlr.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)

	if err := hndlr.QuestionService.UpdateQuestion(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

// @Tags Question
// @Param id path uint true "id"
// @Router /admin/question/{id} [delete]
func (hndlr *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	entity := model.Question{Model: gorm.Model{ID: uint(id)}}

	if err := hndlr.QuestionService.DeleteQuestion(entity.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
