package handler

import (
	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/adapter/presentation/dto"
	"toktok-backend-v1.0.1/internal/adapter/presentation/utils"
	"toktok-backend-v1.0.1/internal/core/port"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	userHandler := UserHandler{
		userService: userService,
	}

	return &userHandler
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	body := dto.CreateUserRequest{}
	if err := utils.ShouldParse(c, &body); err != nil {
		return err
	}

	_, err := h.userService.CreateUser(c.Context(), body.ToUser())
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
