package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(rg *gin.RouterGroup, h *handler.AuthHandler) {
	auths := rg.Group("/auth")
	{
		auths.GET("/login", h.Login)
		auths.POST("/register", h.Register)
	}
}
