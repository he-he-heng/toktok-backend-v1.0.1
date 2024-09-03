package utils

import (
	"github.com/gofiber/fiber/v2"
	"toktok-backend-v1.0.1/internal/adapter/presentation/validator"
	"toktok-backend-v1.0.1/internal/core/domain"
	"toktok-backend-v1.0.1/pkg/errors"
)

func ShouldParse(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	if err := validator.Get().Validate(out); err != nil {
		return errors.Wrap(domain.ErrBadParam, err)
	}

	return nil
}
