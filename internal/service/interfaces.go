package service

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain"
	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/google/uuid"
)

type UserRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*users.User, error)
	FindAll(ctx context.Context) ([]*users.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, user *users.User) error
}

type ProjectRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	FindByOwnerID(ctx context.Context, ownerId uuid.UUID) ([]*domain.Project, error)
	FindAll(ctx context.Context) ([]*domain.Project, error)
	Save(ctx context.Context, project *domain.Project) error
}
