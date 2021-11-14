package main

import (
	"fmt"

	"github.com/ofsahof/library-management/pkg/database"
	"github.com/ofsahof/library-management/pkg/routes"
	"github.com/ofsahof/library-management/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := utils.ParseConfig("config")
	if err != nil {
		panic(err)
	}

	db := database.Init()
	db.SetupRedis(config.DB.Redis.Url, config.DB.Redis.Reset)
	db.SetupGORM(config.DB.Postgres.Host, config.DB.Postgres.Port, config.DB.Postgres.Name, config.DB.Postgres.User, config.DB.Postgres.Password)

	app := fiber.New()

	routes.RegisterAPIRoutes(app)

	app.Listen(fmt.Sprintf(":%d", config.App.Port))
}
