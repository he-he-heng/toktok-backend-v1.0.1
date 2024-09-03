package port

import (
	"context"

	"toktok-backend-v1.0.1/internal/core/domain"
)

type AvatarRepository interface {
	GetAvatar(ctx context.Context, id uint) (*domain.Avatar, error)
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	DeleteAvatar(ctx context.Context, id uint) error
}

type AvatarService interface {
	GetAvatar(ctx context.Context, id uint) (*domain.Avatar, error)
	CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error)
	DeleteAvatar(ctx context.Context, id uint) error
}
