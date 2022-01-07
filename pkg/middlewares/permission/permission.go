package permission

import (
	"github.com/efectn/library-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func New(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := utils.Authority{}

		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		id := int(claims["fields"].(map[string]interface{})["ID"].(float64))
		perm, err := auth.CheckPermission(id, name)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		if perm {
			return c.Next()
		}

		return c.SendStatus(fiber.StatusUnauthorized)
	}
}
