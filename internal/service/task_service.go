package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/taskforge/internal/domain/tasks"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrInternalS      = errors.New("internal server error")
	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskNotCreated = errors.New("tasks cannot be created")
)

type TaskService struct {
	projectRepo ProjectRepo
	taskRepo    TaskRepo
	userRepo    UserRepo
}

func NewTaskService(projectRepo ProjectRepo, taskRepo TaskRepo, userRepo UserRepo) *TaskService {
	return &TaskService{projectRepo: projectRepo, taskRepo: taskRepo, userRepo: userRepo}
}

func (s *TaskService) FindAll(ctx context.Context) ([]*tasks.Task, error) {
	tasks, err := s.taskRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternalS
	}

	return tasks, nil
}

func (s *TaskService) FindAllByAssigneeID(ctx context.Context, assigneeID uuid.UUID) ([]*tasks.Task, error) {
	tasks, err := s.taskRepo.FindAllByAssigneeID(ctx, assigneeID)

	if err != nil {
		return nil, ErrInternalS
	}

	return tasks, nil
}

func (s *TaskService) FindAllByProjectID(ctx context.Context, projectID uuid.UUID) ([]*tasks.Task, error) {
	tasks, err := s.taskRepo.FindAllByProjectID(ctx, projectID)

	if err != nil {
		return nil, ErrInternalS
	}

	return tasks, nil
}

func (s *TaskService) FindByID(ctx context.Context, id uuid.UUID) (*tasks.Task, error) {
	task, err := s.taskRepo.FindByID(ctx, id)

	if err != nil {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (s *TaskService) Create(ctx context.Context, input dto.TaskCreateInput) error {
	_, err := s.projectRepo.FindByID(ctx, input.ProjectID)

	if err != nil {
		return ErrProjectNotFound
	}

	if input.AssigneeID != nil {
		_, err = s.userRepo.FindById(ctx, *input.AssigneeID)
		if err != nil {
			return ErrUserNotFound
		}
	}

	task := &tasks.Task{
		ProjectID:   input.ProjectID,
		Title:       input.Title,
		Description: input.Description,
		Status:      tasks.TodoStatus,
		Priority:    input.Priority,
		AssigneeID:  input.AssigneeID,
		DueDate:     input.DueDate,
	}

	if err := s.taskRepo.Save(ctx, task); err != nil {
		return ErrTaskNotCreated
	}

	return nil
}
