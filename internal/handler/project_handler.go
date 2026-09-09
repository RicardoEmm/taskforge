package handler

import (
	"net/http"

	"github.com/RicardoEmm/taskforge/internal/domain"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) FindAll(c *gin.Context) {
	projects, err := h.projectService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func (h *ProjectHandler) FindByID(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	project, err := h.projectService.FindByID(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func (h *ProjectHandler) FindByOwnerID(c *gin.Context) {
	ownerParsedID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	projects, err := h.projectService.FindByOwnerID(c.Request.Context(), ownerParsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var req dto.ProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.projectService.Create(c.Request.Context(), dto.ProjectCreateInput{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     req.OwnerID,
		Status:      domain.ProjectStatus(req.Status),
	}); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "project was created"})
}
