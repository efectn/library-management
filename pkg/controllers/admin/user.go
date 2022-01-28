package admin

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func (UserController) Index(c *fiber.Ctx) error {
	return c.SendString("index")
}

func (UserController) Store(c *fiber.Ctx) error {
	return c.SendString("store")
}

func (UserController) Show(c *fiber.Ctx) error {
	return c.SendString("show")
}

func (UserController) Update(c *fiber.Ctx) error {
	return c.SendString("update")
}

func (UserController) Destroy(c *fiber.Ctx) error {
	return c.SendString("destroy")
}
