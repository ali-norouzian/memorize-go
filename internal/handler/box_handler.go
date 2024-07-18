package handler

import (
	"memorize/internal/service/box"
	"memorize/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BoxHandler struct {
	BoxService *box.BoxService
	Validator  *validator.Validate
}

func NewBoxHandler(boxService *box.BoxService,
	validator *validator.Validate) *BoxHandler {
	return &BoxHandler{BoxService: boxService,
		Validator: validator,
	}
}

func (hndlr *BoxHandler) SetRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	boxRoutes := routerGroup.Group("/box")
	{
		boxRoutes.GET("/getstatus", hndlr.GetUserBoxesStatus)
	}

	return boxRoutes
}

// @Tags Box
// @Router /box/getstatus [get]
func (hndlr *BoxHandler) GetUserBoxesStatus(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	// 	return
	// }
	userId, _ := c.Get(security.UserIdContextKey)

	entity, err := hndlr.BoxService.GetUserBoxesStatus(userId.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	c.JSON(http.StatusOK, entity)
}
