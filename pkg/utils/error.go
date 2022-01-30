package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ReturnError(c *fiber.Ctx, err interface{}, statusCodeOpt ...int) error {
	statusCode := fiber.StatusForbidden
	if len(statusCodeOpt) > 0 {
		statusCode = statusCodeOpt[0]
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  false,
		"message": err,
	})
}
