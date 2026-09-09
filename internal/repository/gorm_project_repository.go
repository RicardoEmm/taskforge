package repository

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain/projects"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormProjectRepositoy struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) *GormProjectRepositoy {
	return &GormProjectRepositoy{db: db}
}

func (r *GormProjectRepositoy) FindAll(ctx context.Context) ([]*projects.Project, error) {
	var projects []*projects.Project

	if err := r.db.WithContext(ctx).Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *GormProjectRepositoy) FindByID(ctx context.Context, id uuid.UUID) (*projects.Project, error) {
	var project *projects.Project

	if err := r.db.WithContext(ctx).First("id = ?", id).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (r *GormProjectRepositoy) FindByOwnerID(ctx context.Context, ownerId uuid.UUID) ([]*projects.Project, error) {
	var projects []*projects.Project

	if err := r.db.WithContext(ctx).
		Where("owner_id = ?", ownerId).
		Find(&projects).
		Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *GormProjectRepositoy) Save(ctx context.Context, project *projects.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

var _ service.ProjectRepo = (*GormProjectRepositoy)(nil)
