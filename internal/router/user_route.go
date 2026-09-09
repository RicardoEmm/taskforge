package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	users := rg.Group("/users")
	{
		users.GET("", h.FindAll)
		users.GET("/:id", h.FindById)
		users.POST("", h.Create)
	}
}
