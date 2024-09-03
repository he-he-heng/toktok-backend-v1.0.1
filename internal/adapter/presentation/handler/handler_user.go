package handler

import (
	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/adapter/presentation/dto"
	"toktok-backend-v1.0.1/internal/adapter/presentation/utils"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/internal/core/port"
	"toktok-backend-v1.0.1/pkg/errors"
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

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return errors.Wrap(domain.ErrBadParam, "id must be greater than 0")
	}

	user, err := h.userService.GetUser(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetUserResponseOf(user))
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	body := dto.UpdateUserRequest{}
	if err := utils.ShouldParse(c, &body); err != nil {
		return err
	}

	_, err := h.userService.UpdateUser(c.Context(), body.ToUser(), body.ConfirmPassword)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return errors.Wrap(domain.ErrBadParam, "id must be greater than 0")
	}

	err = h.userService.DeleteUser(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
