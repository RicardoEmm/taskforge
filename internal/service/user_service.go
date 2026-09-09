package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/taskforge/internal/domain"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrInternalServer         = errors.New("internal serverver error")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrUserNotCreated         = errors.New("user cannot be created")
)

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := s.userRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternalServer
	}

	return users, nil
}

func (s *UserService) FindById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.FindById(ctx, id)

	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, input dto.UserCreateInput) error {
	exists, err := s.userRepo.ExistsByEmail(ctx, input.Email)

	if err != nil {
		return ErrInternalServer
	}

	if exists {
		return ErrUserEmailAlreadyExists
	}

	if err := s.userRepo.Save(ctx, &domain.User{
		FullName: input.FullName,
		Email:    input.Email,
		Role:     input.Role,
	}); err != nil {
		return ErrUserNotCreated
	}

	return nil
}
