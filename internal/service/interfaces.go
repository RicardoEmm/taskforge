package service

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain"
	"github.com/google/uuid"
)

type UserRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, user *domain.User) error
}
