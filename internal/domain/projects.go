package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "ACTIVE"
	ProjectStatusArchived ProjectStatus = "ARCHIVED"
)

func (s ProjectStatus) IsValid() bool {
	switch s {
	case ProjectStatusActive, ProjectStatusArchived:
		return true
	}
	return false
}

type Project struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string        `gorm:"size:150;not null" json:"name"`
	Description string        `gorm:"type:text" json:"description,omitempty"`
	OwnerID     uuid.UUID     `gorm:"type:uuid;not null;index" json:"owner_id"`
	Owner       *User         `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE" json:"owner,omitempty"`
	Status      ProjectStatus `gorm:"size:20;not null;default:'ACTIVE'" json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `gorm:"index" json:"-"`
}

func (Project) TableName() string {
	return "projects"
}

func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (p *Project) BeforeSave(tx *gorm.DB) error {
	if !p.Status.IsValid() {
		return gorm.ErrInvalidField
	}
	return nil
}
