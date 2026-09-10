package service

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain/projects"
	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/google/uuid"
)

type UserRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*users.User, error)
	FindByEmail(ctx context.Context, email string) (*users.User, error)
	FindAll(ctx context.Context) ([]*users.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, user *users.User) error
}

type ProjectRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*projects.Project, error)
	FindByOwnerID(ctx context.Context, ownerId uuid.UUID) ([]*projects.Project, error)
	FindAll(ctx context.Context) ([]*projects.Project, error)
	Save(ctx context.Context, project *projects.Project) error
}

type TaskRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error)
	FindAll(ctx context.Context) ([]*tasks.Task, error)
	FindAllByProjectID(ctx context.Context, projectID uuid.UUID) ([]*tasks.Task, error)
	FindAllByAssigneeID(ctx context.Context, assigneeID uuid.UUID) ([]*tasks.Task, error)
	Save(ctx context.Context, task *tasks.Task) error
}
