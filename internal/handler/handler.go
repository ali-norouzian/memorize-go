package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	SetRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup
}
