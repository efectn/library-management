package main

import (
	"fmt"
	"ofsahof/library-management/database"
	"ofsahof/library-management/pkg/types"
	"ofsahof/library-management/pkg/utils"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
)

var config types.Config

func main() {
	_, err := toml.DecodeFile("./config/config.toml", &config)
	utils.Check(err)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !!")
	})

	fmt.Println(config)

	db := database.Init()

	db.SetupRedis(config.DB.Redis.Url, config.DB.Redis.Reset)

	db.Redis.Set("d", []byte("d"), 1000)

	fmt.Print(db.Redis.Get("d"))

	app.Listen(":8080")
}
