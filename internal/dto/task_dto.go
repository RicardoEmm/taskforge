package dto

import (
	"time"

	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/google/uuid"
)

type TaskCreateInput struct {
	ProjectID   uuid.UUID
	Title       string
	Description *string
	Status      tasks.TaskStatus
	Priority    tasks.TaskPriority
	AssigneeID  *uuid.UUID
	DueDate     *time.Time
}

type CreateTaskRequest struct {
	ProjectID   string     `json:"project_id" binding:"required,uuid"`
	Title       string     `json:"title" binding:"required,max=180"`
	Description *string    `json:"description" binding:"omitempty,max=2000"`
	Status      string     `json:"status" binding:"required,oneof=TODO IN_PROGRES DONE"`
	Priority    string     `json:"priority" binding:"required,oneof=LOW MEDIUM HIGH"`
	AssigneeID  *string    `json:"assignee_id" binding:"omitempty,uuid"`
	DueDate     *time.Time `json:"due_date" binding:"omitempty"`
}
