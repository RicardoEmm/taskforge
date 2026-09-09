package dto

import (
	"github.com/RicardoEmm/taskforge/internal/domain/projects"
	"github.com/google/uuid"
)

type ProjectRequest struct {
	Name        string `json:"name" binding:"required,max=150"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=ACTIVE ARCHIVED"`
}

type ProjectCreateInput struct {
	Name        string
	Description string
	OwnerId     uuid.UUID
	Status      projects.ProjectStatus
}
