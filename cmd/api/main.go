package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/efectn/library-management/pkg/database/ent/runtime"
	"github.com/efectn/library-management/pkg/database/seeds"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/routes"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/efectn/library-management/pkg/webserver"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Init the app
func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		api.App.Shutdown()
		os.Exit(1)
	}()
}

// TODO: Can't access endpoints when prefork was opened.

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

	// Migrate
	err = api.App.DB.MigrateModels()
	if err != nil {
		api.App.Logger.Panic().Err(err).Msg("")
	}

	// Seed
	api.App.DB.SeedModels(api.App.Logger, seeds.PermissionSeeder{}, seeds.RoleSeeder{}, seeds.UserSeeder{})

	// Register Routes & Listen
	routes.RegisterAPIRoutes(api.App.Fiber)

	// Listen the app
	if !fiber.IsChild() {
		err = api.App.Run()
		if err != nil {
			api.App.Logger.Panic().Err(err).Msg("")
		}
	}
}

// Execute app
func main() {
	Execute()
}
