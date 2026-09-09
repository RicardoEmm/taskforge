package repository

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) FindAll(ctx context.Context) ([]*tasks.Task, error) {
	var tasks []*tasks.Task

	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *GormTaskRepository) FindAllByAssigneeID(ctx context.Context, assigneeID uuid.UUID) ([]*tasks.Task, error) {
	var tasks []*tasks.Task

	if err := r.db.WithContext(ctx).
		Where("assignee_id = ?", assigneeID).
		Find(&tasks).
		Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *GormTaskRepository) FindByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	var task *tasks.Task

	if err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *GormTaskRepository) FindAllByProjectID(ctx context.Context, projectID uuid.UUID) ([]*tasks.Task, error) {
	var tasks []*tasks.Task

	if err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).
		Find(&tasks).
		Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *GormTaskRepository) Save(ctx context.Context, task *tasks.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

var _ service.TaskRepo = (*GormTaskRepository)(nil)
