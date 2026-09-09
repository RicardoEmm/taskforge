package tasks

import (
	"time"

	"github.com/RicardoEmm/taskforge/internal/domain/projects"
	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	ProjectID   uuid.UUID         `gorm:"type:uuid;not null;index" json:"project_id"`
	Project     *projects.Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project,omitempty"`
	Title       string            `gorm:"size:180;not null" json:"title"`
	Description *string           `gorm:"type:text" json:"description,omitempty"`
	Status      TaskStatus        `gorm:"size:20;not null;default:'TODO'" json:"status"`
	Priority    TaskPriority      `gorm:"size:10;not null" json:"priority"`
	AssigneeID  *uuid.UUID        `gorm:"type:uuid;not null" json:"assignee_id,omitempty"`
	Assignee    *users.User       `gorm:"foreignKey:AssigneeID" json:"assignee,omitempty"`
	DueDate     *time.Time        `gorm:"type:date" json:"due_date,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"-"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

func (t *Task) BeforeSave(tx *gorm.DB) error {
	if !t.Status.IsValid() {
		return gorm.ErrInvalidField
	}

	if !t.Priority.IsValid() {
		return gorm.ErrInvalidField
	}

	return nil
}
