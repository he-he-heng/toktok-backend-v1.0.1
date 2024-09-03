package dto

import "toktok-backend-v1.0.1/internal/core/domain"

// POST v1/api/users
type CreateUserRequest struct {
	LoginID  string `json:"login_id,omitempty" validate:"gte=4,lte=18"`
	Password string `json:"password,omitempty" validate:"gte=6,lte=32"`
}

func (c CreateUserRequest) ToUser() *domain.User {
	return &domain.User{
		LoginID:  c.LoginID,
		Password: c.Password,
	}
}

type UpdateUserRequest struct {
	LoginID         string  `json:"login_id" validate:"omitempty,gte=4,lte=18"`
	Password        string  `json:"password" validate:"omitempty,gte=6,lte=32"`
	Email           *string `json:"email" validate:"omitempty,email"`
	ConfirmPassword string  `json:"confirm_password" validate:"omitempty,gte=6,lte=32"`
}

func (u UpdateUserRequest) ToUser() *domain.User {
	return &domain.User{
		LoginID:  u.LoginID,
		Password: u.Password,
		Email:    u.Email,
	}
}
