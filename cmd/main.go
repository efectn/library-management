package main

import (
	"github.com/efectn/library-management/pkg/app"
	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/database/seeds"
	"github.com/efectn/library-management/pkg/routes"
	"github.com/efectn/library-management/pkg/utils/config"
)

func main() {
	// Parse Config
	config, err := config.ParseConfig("config")
	if err != nil {
		panic(err)
	}

	// Init App
	app.App = app.New(config)

	// Database
	err = app.App.SetupDB()
	if err != nil {
		panic(err)
	}

	err = app.App.DB.MigrateModels(&models.Users{}, &models.Role{}, &models.Permission{})
	if err != nil {
		panic(err)
	}

	app.App.DB.SeedModels(seeds.UserSeeder{})

	// Register Routes & Listen
	routes.RegisterAPIRoutes(app.App.Fiber, config)

	err = app.App.Listen()
	if err != nil {
		panic(err)
	}
}
