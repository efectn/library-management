package main

import (
	"github.com/efectn/library-management/pkg/app"
	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/database/seeds"
	"github.com/efectn/library-management/pkg/globals"
	"github.com/efectn/library-management/pkg/routes"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/rs/zerolog/log"
)

// Fix:
// - Prefork not working with zerolog.
func main() {
	// Parse Config
	config, err := config.ParseConfig("config")
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	// Init App
	globals.App = app.New(config)

	// Logger
	globals.App.SetupLogger()

	// Database
	err = globals.App.SetupDB()
	if err != nil {
		globals.App.Logger.Panic().Err(err).Msg("")
	}

	// Migrate
	err = globals.App.DB.MigrateModels(&models.Users{}, &models.Role{}, &models.Permission{})
	if err != nil {
		globals.App.Logger.Panic().Err(err).Msg("")
	}

	// Seed
	globals.App.DB.SeedModels(seeds.PermissionSeeder{}, seeds.RoleSeeder{}, seeds.UserSeeder{})

	// Register Routes & Listen
	routes.RegisterAPIRoutes(globals.App.Fiber, config)

	// Listen the app
	err = globals.App.Listen()
	if err != nil {
		globals.App.Logger.Panic().Err(err).Msg("")
	}
}
