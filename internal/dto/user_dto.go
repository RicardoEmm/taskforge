package dto

import "github.com/RicardoEmm/taskforge/internal/domain"

type UserRequest struct {
	FullName string `json:"full_name" binding:"required,max=120"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required,oneof=ADMIN MEMBER"`
}

type UserCreateInput struct {
	FullName string
	Email    string
	Role     domain.UserRole
}
