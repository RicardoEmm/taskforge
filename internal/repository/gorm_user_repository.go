package repository

import (
	"context"

	"github.com/RicardoEmm/taskforge/internal/domain"
	"github.com/RicardoEmm/taskforge/internal/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Table("users").
		Where("email = ?", email).
		Count(&count).
		Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *GormUserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *GormUserRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user *domain.User

	if err := r.db.WithContext(ctx).First("id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *GormUserRepository) Save(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

var _ service.UserRepo = (*GormUserRepository)(nil)
