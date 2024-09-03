package repository

import (
	"context"

	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql"
	"toktok-backend-v1.0.1/internal/adapter/persistence/mysql/utils"
	"toktok-backend-v1.0.1/internal/core/domain"
)

type UserRepository struct {
	db *mysql.Database
}

func NewUserRepository(database *mysql.Database) *UserRepository {
	userRepository := UserRepository{
		db: database,
	}

	return &userRepository
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return user, err
}

func (r *UserRepository) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	queryUser := domain.User{}
	err := r.db.WithContext(ctx).First(&queryUser, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &queryUser, err
}

func (r *UserRepository) GetUserByLoginID(ctx context.Context, loginID string) (*domain.User, error) {
	queriedUser := domain.User{}

	err := r.db.WithContext(ctx).Where("login_id = ?", loginID).Find(&queriedUser).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &queriedUser, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&domain.User{}, id).Error
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
