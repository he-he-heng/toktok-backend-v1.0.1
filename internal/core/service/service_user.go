package service

import (
	"context"

	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/internal/core/port"
	"toktok-backend-v1.0.1/internal/core/service/encryption"
	"toktok-backend-v1.0.1/pkg/errors"
)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	userService := UserService{
		userRepository: userRepository,
	}

	return &userService
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := encryption.HashPassword(user.Password)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}

	user.Password = hashedPassword
	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil

}

func (s *UserService) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	gotUser, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}
func (s *UserService) UpdateUser(ctx context.Context, user *domain.User, confirmPassword string) (*domain.User, error) {

	if user.LoginID != "" || user.Password != "" {
		gotUser, err := s.userRepository.GetUser(ctx, user.ID)
		if err != nil {
			return nil, err
		}

		err = encryption.VerifyPassword(confirmPassword, gotUser.Password)
		if err != nil {
			return nil, errors.Wrap(domain.ErrUnauthorized, err)
		}

		hashPassword, err := encryption.HashPassword(user.Password)
		if err != nil {
			return nil, errors.Wrap(domain.ErrInternalServerError, err)
		}

		user.Password = hashPassword
	}

	updatedUser, err := s.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil

}
