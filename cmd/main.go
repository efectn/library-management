package main

import (
	"fmt"

	"github.com/efectn/library-management/pkg/database"
	"github.com/efectn/library-management/pkg/database/seeds"
	"github.com/efectn/library-management/pkg/routes"
	"github.com/efectn/library-management/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := utils.ParseConfig("config")
	if err != nil {
		panic(err)
	}
	utils.Config = config

	database.DB.SetupRedis(config.DB.Redis.Url, config.DB.Redis.Reset)
	database.DB.SetupGORM(config.DB.Postgres.Host, config.DB.Postgres.Port, config.DB.Postgres.Name, config.DB.Postgres.User, config.DB.Postgres.Password)
	database.DB.MigrateModels()
	seeds.SeedModels(seeds.UserSeeder{})

	app := fiber.New()

	routes.RegisterAPIRoutes(app, config)

	app.Listen(fmt.Sprintf(":%d", config.App.Port))
}
