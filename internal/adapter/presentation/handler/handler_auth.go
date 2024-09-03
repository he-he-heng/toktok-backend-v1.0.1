package handler

import (
	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/adapter/presentation/dto"
	"toktok-backend-v1.0.1/internal/adapter/presentation/middleware"
	"toktok-backend-v1.0.1/internal/adapter/presentation/utils"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/internal/core/port"
	"toktok-backend-v1.0.1/pkg/errors"
)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	body := dto.LoginRequest{}
	if err := utils.ShouldParse(c, &body); err != nil {
		return err
	}

	access, refresh, err := h.authService.Login(c.Context(), body.LoginID, body.Password)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.LoginResponseOf(access, refresh))
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	body := dto.RefreshRequest{}
	if err := utils.ShouldParse(c, &body); err != nil {
		return err
	}

	access, err := h.authService.Refresh(c.Context(), body.AccessToken)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.RefreshResponseOf(access))
}

func (h *AuthHandler) Validation(c *fiber.Ctx) error {
	// next from GuardMiddleware.TokenValidate
	tokenPayload, ok := c.Locals(middleware.AuthorizationPayloadKey).(*domain.TokenPayload)
	if !ok {
		return errors.Wrap(domain.ErrUnauthorized, "invalid access token")
	}

	return c.Status(fiber.StatusOK).JSON(dto.ValidateResponseOf(tokenPayload))
}
