package webserver

import (
	"os"
	"runtime"

	"github.com/efectn/library-management/pkg/database"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppSkel struct {
	Fiber  *fiber.App
	DB     *database.Database
	Config *config.Config
	Logger zerolog.Logger
}

type PreforkHook struct{}

func New(configPart *config.Config) *AppSkel {
	app := &AppSkel{
		Fiber: fiber.New(fiber.Config{
			AppName:               configPart.App.Name,
			ServerHeader:          configPart.App.Name,
			Prefork:               configPart.App.Prefork,
			DisableStartupMessage: true,
		}),
		DB:     database.Init(),
		Config: configPart,
		Logger: zerolog.Logger{},
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

func (app *AppSkel) SetupLogger() error {
	zerolog.TimeFieldFormat = app.Config.Logger.TimeFormat

	if app.Config.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(app.Config.Logger.Level)

	app.Logger = log.Hook(PreforkHook{})

	return nil
}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}

func (app *AppSkel) Run() error {
	// Custom Startup Messages
	host, port := config.ParseAddr(app.Config.App.Port)
	if host == "" {
		if app.Fiber.Config().Network == "tcp6" {
			host = "[::1]"
		} else {
			host = "0.0.0.0"
		}
	}

	// ASCII Art
	app.Logger.Info().Msg("█████       ██████   ██████  █████████   ")
	app.Logger.Info().Msg("░░███       ░░██████ ██████  ███░░░░░███ ")
	app.Logger.Info().Msg(" ░███        ░███░█████░███ ░███    ░░░  ")
	app.Logger.Info().Msg(" ░███        ░███░░███ ░███ ░░█████████  ")
	app.Logger.Info().Msg(" ░███        ░███ ░░░  ░███  ░░░░░░░░███ ")
	app.Logger.Info().Msg(" ░███      █ ░███      ░███  ███    ░███ ")
	app.Logger.Info().Msg(" ███████████ █████     █████░░█████████  ")
	app.Logger.Info().Msg("░░░░░░░░░░░ ░░░░░     ░░░░░  ░░░░░░░░░   ")

	// Information message
	app.Logger.Info().Msg(app.Fiber.Config().AppName + " is running at the moment!")

	// Debug informations
	if !app.Config.App.Production {
		prefork := "Enabled"
		procs := runtime.GOMAXPROCS(0)
		if !app.Config.App.Prefork {
			procs = 1
			prefork = "Disabled"
		}

		app.Logger.Debug().Msgf("Host: %s", host)
		app.Logger.Debug().Msgf("Port: %s", port)
		app.Logger.Debug().Msgf("Prefork: %s", prefork)
                app.Logger.Debug().Msgf("Handlers: %s", app.Fiber.HandlersCount())
		app.Logger.Debug().Msgf("Processes: %d", procs)
		app.Logger.Debug().Msgf("PID: %d", os.Getpid())
	}

	// Listen the app
	err := app.Fiber.Listen(app.Config.App.Port)
	if err != nil {
		return err
	}

	return nil
}
