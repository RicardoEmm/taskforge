package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/RicardoEmm/taskforge/internal/middleware"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AuthHandler    *handler.AuthHandler
	ProjectHandler *handler.ProjectHandler
	TaskHandler    *handler.TaskHandler
	UserHandler    *handler.UserHandler
}

func Setup(h Handlers, tokenService *service.TokenService) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		registerAuthRoutes(api, h.AuthHandler)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired(tokenService))
		{
			registerUserRoutes(api, h.UserHandler)
			registerTaskRoutes(api, h.TaskHandler)
			registerProductRoutes(api, h.ProjectHandler)
		}
	}

	return router
}
