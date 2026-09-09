package handler

import (
	"net/http"

	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/RicardoEmm/taskforge/internal/httputil"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) FindAll(c *gin.Context) {
	tasks, err := h.taskService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) FindAllByAssigneeID(c *gin.Context) {

	assigneeID, ok := httputil.ParseUUID(c, c.Param("id"), "assignee_id")

	if !ok {
		return
	}

	tasks, err := h.taskService.FindAllByAssigneeID(c.Request.Context(), assigneeID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) FindAllByProjectID(c *gin.Context) {

	projectID, ok := httputil.ParseUUID(c, c.Param("id"), "project_id")

	if !ok {
		return
	}

	tasks, err := h.taskService.FindAllByProjectID(c.Request.Context(), projectID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) FindByID(c *gin.Context) {

	parsedID, ok := httputil.ParseUUID(c, c.Param("id"), "id")

	if !ok {
		return
	}

	task, err := h.taskService.FindByID(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectID, ok := httputil.ParseUUID(c, c.Param("id"), "project_id")

	if !ok {
		return
	}

	var assigneeParsedID *uuid.UUID

	if req.AssigneeID != nil {
		parsedID, err := uuid.Parse(*req.AssigneeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid assignee id"})
			return
		}
		assigneeParsedID = &parsedID
	}

	input := dto.TaskCreateInput{
		ProjectID:   projectID,
		Title:       req.Title,
		Description: req.Description,
		Status:      tasks.TaskStatus(req.Status),
		Priority:    tasks.TaskPriority(req.Priority),
		AssigneeID:  assigneeParsedID,
		DueDate:     req.DueDate,
	}

	if err := h.taskService.Create(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "task was created"})
}
