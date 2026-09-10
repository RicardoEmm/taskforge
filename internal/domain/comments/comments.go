package comments

import (
	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	TaskID  uuid.UUID   `gorm:"type:uuid;not null" json:"task_id"`
	Task    *tasks.Task `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"task,omitempty"`
	UserID  uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	User    *users.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Content string      `gorm:"type:text;not null" json:"content"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
