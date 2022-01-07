package main

import (
	"github.com/efectn/library-management/pkg/database/seeds"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/routes"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/efectn/library-management/pkg/webserver"
	"github.com/rs/zerolog/log"
)

// Fix:
// - Prefork not working with zerolog.

// Execute the app
func Execute() {
	// Parse Config
	parseConfig, err := config.ParseConfig("api")
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	// Init App
	api.App = webserver.New(parseConfig)

	// Logger
	err = api.App.SetupLogger()
	if err != nil {
		api.App.Logger.Panic().Err(err).Msg("")
	}

	// Database
	err = api.App.SetupDB()
	if err != nil {
		api.App.Logger.Panic().Err(err).Msg("")
	}

	//db := api.App.DB.Ent
	//db.Permission.Create().SetName("test")
	//db.Permission.Create().SetName("test-2")
	//db.Role.Create().SetName("role")
	//db.User.Create().AD

	// Migrate
	err = api.App.DB.MigrateModels()
	if err != nil {
		api.App.Logger.Panic().Err(err).Msg("")
	}

	// Seed
	api.App.DB.SeedModels(seeds.PermissionSeeder{}, seeds.RoleSeeder{}, seeds.UserSeeder{})

	// Register Routes & Listen
	routes.RegisterAPIRoutes(api.App.Fiber)

	// Listen the app
	err = api.App.Run()
	if err != nil {
		api.App.Logger.Panic().Err(err).Msg("")
	}
}

// Execute app
func main() {
	Execute()
}
