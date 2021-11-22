package app

import (
	"fmt"

	"github.com/efectn/library-management/pkg/database"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type AppSkel struct {
	Fiber  *fiber.App
	DB     *database.Database
	Config *config.Config
}

var App *AppSkel

func New(configPart *config.Config) *AppSkel {
	app := &AppSkel{
		Fiber:  fiber.New(),
		DB:     database.Init(),
		Config: configPart,
	}

	// Register some middlewares
	app.Fiber.Use(compress.New(compress.Config{
		Next:  config.IsEnabled(configPart.Middleware.Compress.Enable),
		Level: configPart.Middleware.Compress.Level,
	}))

	app.Fiber.Use(recover.New(recover.Config{
		Next: config.IsEnabled(configPart.Middleware.Recover.Enable),
	}))

	return app
}

func (app *AppSkel) SetupDB() error {
	err := app.DB.SetupRedis(app.Config.DB.Redis.Url, app.Config.DB.Redis.Reset)
	if err != nil {
		return err
	}

	err = app.DB.SetupGORM(app.Config.DB.Postgres.Host, app.Config.DB.Postgres.Port, app.Config.DB.Postgres.Name, app.Config.DB.Postgres.User, app.Config.DB.Postgres.Password)
	if err != nil {
		return err
	}

	return nil
}

func (app *AppSkel) Listen() error {
	err := app.Fiber.Listen(fmt.Sprintf(":%d", app.Config.App.Port))
	if err != nil {
		return err
	}

	return nil
}
