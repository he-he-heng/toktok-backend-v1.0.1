package service

import (
	"context"

	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/internal/core/port"
	"toktok-backend-v1.0.1/internal/core/service/encryption"
	"toktok-backend-v1.0.1/pkg/errors"
)

type AuthService struct {
	userRepository port.UserRepository
	tokenService   port.TokenService
}

func NewAuthService(userRepository port.UserRepository, tokenService port.TokenService) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}
func (s *AuthService) Login(ctx context.Context, loginID, password string) (string, string, error) {
	queriedUser, err := s.userRepository.GetUserByLoginID(ctx, loginID)
	if err != nil {
		return "", "", err
	}

	err = encryption.VerifyPassword(password, queriedUser.Password)
	if err != nil {
		return "", "", errors.Wrap(domain.ErrUnauthorized, err)
	}

	accessToken, err := s.tokenService.CreateToken(domain.AccessToken, queriedUser)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenService.CreateToken(domain.RefreshToken, queriedUser)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (accessToekn string, err error) {
	tokenPayload, err := s.tokenService.VerifyToken(token)
	if err != nil {
		return "", err
	}

	if tokenPayload.TokenType != domain.RefreshToken {
		return "", errors.Wrap(domain.ErrUnauthorized, "token type is not refresh token")
	}

	gotUser, err := s.userRepository.GetUser(ctx, uint(tokenPayload.Iss))
	if err != nil {
		return "", err
	}

	accessToken, err := s.tokenService.CreateToken(domain.AccessToken, gotUser)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
