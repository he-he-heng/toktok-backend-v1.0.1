package port

import (
	"context"

	"toktok-backend-v1.0.1/internal/core/domain"
)

type TokenService interface {

	// CreateToken 함수는 tokenPayload와 user값을 기반으로 새로운 토큰을 생성합니다.
	CreateToken(tokenType domain.TokenPayloadTokenType, user *domain.User) (string, error)

	// VerifyToken 함수는 token을 검증하여 tokenPayload로 변환해줍니다.
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	Login(ctx context.Context, loginID, password string) (string, string, error)

	Refresh(ctx context.Context, token string) (accessToken string, err error)
}
