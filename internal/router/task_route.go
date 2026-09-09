package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(rg *gin.RouterGroup, h *handler.TaskHandler) {
	tasks := rg.Group("/tasks")
	{
		tasks.GET("/:id", h.FindByID)
		tasks.GET("/", h.FindAll)
		tasks.GET("/assignee/:id", h.FindAllByAssigneeID)
		tasks.GET("/project/:id", h.FindAllByProjectID)
		tasks.POST("/", h.Create)
	}
}
