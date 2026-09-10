package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FullName     string    `gorm:"size:120;not null" json:"full_name"`
	Email        string    `gorm:"size:160;not null;uniqueIndex" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         UserRole  `gorm:"size:20;not null;default:MEMBER" json:"role"`
	AvatarPath   *string   `gorm:"size:255" json:"avatar_path"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if !u.Role.IsValid() {
		return gorm.ErrInvalidField
	}
	return nil
}
