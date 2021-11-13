package main

import (
	"fmt"

	"github.com/ofsahof/library-management/pkg/database"
	"github.com/ofsahof/library-management/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := utils.ParseConfig("config")
	utils.Check(err)

	db := database.Init()
	db.SetupRedis(config.DB.Redis.Url, config.DB.Redis.Reset)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !!")
	})

	app.Listen(fmt.Sprintf(":%d", config.App.Port))
}
