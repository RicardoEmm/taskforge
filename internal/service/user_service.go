package service

import (
	"context"
	"errors"

	"github.com/RicardoEmm/taskforge/internal/domain/users"
	"github.com/RicardoEmm/taskforge/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrInternalServer         = errors.New("internal serverver error")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrUserNotCreated         = errors.New("user cannot be created")
	ErrUserBadCredentials     = errors.New("bad credentials")
)

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindAll(ctx context.Context) ([]*users.User, error) {
	users, err := s.userRepo.FindAll(ctx)

	if err != nil {
		return nil, ErrInternalServer
	}

	return users, nil
}

func (s *UserService) FindById(ctx context.Context, id uuid.UUID) (*users.User, error) {
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

	if err := s.userRepo.Save(ctx, &users.User{
		FullName: input.FullName,
		Email:    input.Email,
		Role:     input.Role,
	}); err != nil {
		return ErrUserNotCreated
	}

	return nil
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*users.User, error) {
	exists, err := s.userRepo.ExistsByEmail(ctx, email)

	if err != nil {
		return nil, ErrInternalServer
	}

	if exists {
		return nil, ErrUserEmailAlreadyExists
	}

	hash, err := HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := &users.User{
		FullName:     name,
		Email:        email,
		PasswordHash: hash,
		Role:         users.MemberRole,
	}

	if err := s.userRepo.Save(ctx, user); err != nil {
		return nil, ErrUserNotCreated
	}

	return user, nil
}

func (s *UserService) ValidateCredentials(ctx context.Context, email, password string) (*users.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)

	if err != nil {
		return nil, ErrUserBadCredentials
	}

	if !CheckPassword(user.PasswordHash, password) {
		return nil, ErrUserBadCredentials
	}

	return user, nil
}
