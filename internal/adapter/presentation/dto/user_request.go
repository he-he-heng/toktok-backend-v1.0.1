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

// GET /v1/api/users/{id}
