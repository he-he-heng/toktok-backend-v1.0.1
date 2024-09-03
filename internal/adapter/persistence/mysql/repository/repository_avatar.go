package repository

import (
	"context"

	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql"
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql/utils"
	"toktok-backend-v1.0.1/internal/core/domain"
)

type AvatarRepository struct {
	db *mysql.Database
}

func (r *AvatarRepository) GetAvatar(ctx context.Context, id uint) (*domain.Avatar, error) {
	queryAvatar := domain.Avatar{}
	err := r.db.WithContext(ctx).First(&queryAvatar, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &queryAvatar, nil
}

func (r *AvatarRepository) CreateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	err := r.db.WithContext(ctx).Create(avatar).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return avatar, nil
}

func (r *AvatarRepository) UpdateAvatar(ctx context.Context, avatar *domain.Avatar) (*domain.Avatar, error) {
	err := r.db.WithContext(ctx).Save(avatar).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return avatar, nil
}

func (r *AvatarRepository) DeleteAvatar(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&domain.Avatar{}, id).Error
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
