package port

import (
	"context"

	"toktok-backend-v1.0.1/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(cxx context.Context, user *domain.User, confirmPassword string) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}
