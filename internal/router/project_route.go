package router

import (
	"github.com/RicardoEmm/taskforge/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.RouterGroup, h *handler.ProjectHandler) {
	products := rg.Group("/products")
	{
		products.GET("", h.FindAll)
		products.GET("/:id", h.FindByID)
		products.GET("/owner/:id", h.FindByOwnerID)
		products.POST("", h.Create)
	}
}
