package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/internal/core/port"
	"toktok-backend-v1.0.1/pkg/errors"
)

const (
	AuthorizationPayloadKey = "authorization_payload"
)

type GuardMiddleware struct {
	tokenService port.TokenService
}

func NewGuardMiddlware(tokenService port.TokenService) *GuardMiddleware {
	return &GuardMiddleware{tokenService: tokenService}
}

func (m *GuardMiddleware) TokenValidate(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	payload, err := m.tokenValidate(header)
	if err != nil {
		return err
	}

	c.Locals(AuthorizationPayloadKey, payload)
	return c.Next()
}

func (m *GuardMiddleware) FilterValidAccess(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	payload, err := m.tokenValidate(header)
	if err != nil {
		return err
	}

	// /v1/api/users/:id --> 여기 id값을 가져온다.
	id, err := c.ParamsInt("id", 0)
	if err != nil || id < 0 {
		return errors.Wrap(domain.ErrBadParam, "id must be greater than or equal to 0")
	}

	if payload.Iss != id {
		return errors.Wrap(domain.ErrUnauthorized, "invalid access token")
	}

	return c.Next()
}

func (m *GuardMiddleware) tokenValidate(header string) (*domain.TokenPayload, error) {

	// Authorization 헤더의 필드값이 있는가?
	if len(header) == 0 {
		return nil, errors.Wrap(domain.ErrUnauthorized, "err empty authorization header")
	}

	// 잘못된 필드값임을 알림
	fields := strings.Fields(header)
	if len(fields) != 2 {
		return nil, errors.Wrap(domain.ErrUnauthorized, "err empty authorization header")

	}

	// bearer 헤더인지 확인
	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, errors.Wrap(domain.ErrUnauthorized, "err invalid authorization header")

	}

	tokenString := fields[1]
	payload, err := m.tokenService.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	if payload.TokenType == domain.RefreshToken {
		return nil, errors.Wrap(domain.ErrUnauthorized, "token type not match")
	}

	return payload, nil

}
