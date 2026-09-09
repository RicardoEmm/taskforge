package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductHandler *handler.ProjectHandler
	TaskHandler    *handler.TaskHandler
	UserHandler    *handler.UserHandler
}

func Setup(h Handlers) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		registerUserRoutes(api, h.UserHandler)
		registerTaskRoutes(api, h.TaskHandler)
		registerProductRoutes(api, h.ProductHandler)
	}

	return router
}
