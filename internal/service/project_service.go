package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/taskforge/internal/domain/projects"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrInternal          = errors.New("internal server error")
	ErrProjectNotFounf   = errors.New("project not found")
	ErrProjectNotCreated = errors.New("project cannot be created")
)

type ProjectService struct {
	projectRepo ProjectRepo
	userRepo    UserRepo
}

func NewProjectService(projectRepo ProjectRepo, userRepo UserRepo) *ProjectService {
	return &ProjectService{projectRepo: projectRepo, userRepo: userRepo}
}

func (s *ProjectService) FindAll(ctx context.Context) ([]*projects.Project, error) {
	projects, err := s.projectRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternal
	}

	return projects, nil
}

func (s *ProjectService) FindByID(ctx context.Context, id uuid.UUID) (*projects.Project, error) {
	project, err := s.projectRepo.FindByID(ctx, id)

	if err != nil {
		return nil, ErrProjectNotFounf
	}

	return project, nil
}

func (s *ProjectService) FindByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*projects.Project, error) {
	projects, err := s.projectRepo.FindByOwnerID(ctx, ownerID)

	if err != nil {
		return nil, ErrInternal
	}

	return projects, nil
}

func (s *ProjectService) Create(ctx context.Context, input dto.ProjectCreateInput) error {
	owner, err := s.userRepo.FindById(ctx, input.OwnerId)

	if err != nil {
		return ErrUserNotFound
	}

	if err := s.projectRepo.Save(ctx, &projects.Project{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     owner.ID,
	}); err != nil {
		return ErrProjectNotCreated
	}

	return nil
}
