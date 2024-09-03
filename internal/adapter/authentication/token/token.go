package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"toktok-backend-v1.0.1/internal/config"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/pkg/errors"
)

type TokenService struct {
	secretKey []byte

	accessDuration  time.Duration
	refreshDuration time.Duration
}

func NewTokenService(config *config.Config) (*TokenService, error) {
	accessDuration, err := time.ParseDuration(config.Token.AccessDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}

	refreshDuration, err := time.ParseDuration(config.Token.RefreshDuration)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServerError, err)
	}

	return &TokenService{
		secretKey:       []byte(config.Token.Key),
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}, nil
}

func (s *TokenService) CreateToken(tokenType domain.TokenPayloadTokenType, user *domain.User) (string, error) {
	var duration time.Duration = s.accessDuration

	if tokenType == domain.RefreshToken {
		duration = s.refreshDuration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       user.ID,                         // uint
		"exp":       time.Now().Add(duration).Unix(), //int64
		"ita":       time.Now().Unix(),
		"tokenType": tokenType,
		"role":      user.Role,
	})

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", errors.Wrap(domain.ErrInternalServerError, err)
	}

	return tokenString, nil
}

func (s *TokenService) VerifyToken(tokenString string) (*domain.TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(domain.ErrUnauthorized, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return s.secretKey, nil
	})

	if err != nil {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.Wrap(domain.ErrUnauthorized, err)
	}

	return &domain.TokenPayload{
		Iss:       int(claims["iss"].(float64)),
		Exp:       int64(claims["exp"].(float64)),
		Ita:       int64(claims["ita"].(float64)),
		TokenType: domain.TokenPayloadTokenType(claims["tokenType"].(string)),
		Role:      domain.UserRoleType(claims["role"].(string)),
	}, nil

}
