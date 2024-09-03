package sniff

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var CustomErrorHandler = func(c *fiber.Ctx, err error) error {
	msg := err.Error()
	status := fiber.StatusInternalServerError

	gotStatus, originErrMsg, err := errSetGet().Get(err)
	if err == nil {
		status = gotStatus
		msg = fmt.Sprintf("[%s] %s", originErrMsg, msg)
	}

	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"msg":    msg,
	})
}
