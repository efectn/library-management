package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ReturnErrorMessage(c *fiber.Ctx, err error, statusCodeOpt ...int) error {
	statusCode := fiber.StatusForbidden
	if len(statusCodeOpt) > 0 {
		statusCode = statusCodeOpt[0]
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": err.Error(),
	})
}
