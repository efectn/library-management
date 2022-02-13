package webserver

import (
	"crypto/tls"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/dgrr/http2"
	"github.com/efectn/library-management/pkg/database"
	"github.com/efectn/library-management/pkg/globals"
	"github.com/efectn/library-management/pkg/utils/config"
	"github.com/efectn/library-management/pkg/utils/convert"
	"github.com/efectn/library-management/pkg/utils/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppSkel struct {
	Fiber     *fiber.App
	DB        *database.Database
	Config    *config.Config
	Logger    zerolog.Logger
	Validator *validator.Validate
}

type PreforkHook struct{}

func New(configPart *config.Config) *AppSkel {
	app := &AppSkel{
		Fiber: fiber.New(fiber.Config{
			AppName:               configPart.App.Name,
			ServerHeader:          configPart.App.Name,
			Prefork:               configPart.App.Prefork,
			DisableStartupMessage: true,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				var messages interface{}

				if e, ok := err.(*errors.Error); ok {
					code = e.Code
					messages = e.Message
				}

				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
					messages = e.Message
				}

				return c.Status(code).JSON(fiber.Map{
					"status":   false,
					"messages": messages,
				})
			},
			IdleTimeout:       configPart.App.IdleTimeout * time.Second,
			EnablePrintRoutes: configPart.App.PrintRoutes,
		}),
		DB:        database.Init(),
		Config:    configPart,
		Logger:    zerolog.Logger{},
		Validator: validator.New(),
	}

	// Register several middlewares
	app.Fiber.Use(compress.New(compress.Config{
		Next:  config.IsEnabled(configPart.Middleware.Compress.Enable),
		Level: configPart.Middleware.Compress.Level,
	}))

	app.Fiber.Use(recover.New(recover.Config{
		Next: config.IsEnabled(configPart.Middleware.Recover.Enable),
	}))

	app.Fiber.Use(pprof.New(pprof.Config{
		Next: config.IsEnabled(configPart.Middleware.Pprof.Enable),
	}))

	app.Fiber.Get(configPart.Middleware.Monitor.Path, monitor.New(monitor.Config{
		Next: config.IsEnabled(configPart.Middleware.Monitor.Enable),
	}))

	return app
}

func (app *AppSkel) SetupDB() error {
	err := app.DB.SetupRedis(app.Config.DB.Redis.Url, app.Config.DB.Redis.Reset)
	if err != nil {
		return err
	}

	err = app.DB.SetupEnt(app.Config.DB.Postgres.Host,
		app.Config.DB.Postgres.Port,
		app.Config.DB.Postgres.User,
		app.Config.DB.Postgres.Password,
		app.Config.DB.Postgres.Name)
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
	ascii, err := os.ReadFile("./storage/ascii_art.txt")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(convert.UnsafeString(ascii), "\n") {
		app.Logger.Info().Msg(line)
	}

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

		app.Logger.Debug().Msgf("Version: %s", globals.VERSION)
		app.Logger.Debug().Msgf("Host: %s", host)
		app.Logger.Debug().Msgf("Port: %s", port)
		app.Logger.Debug().Msgf("Prefork: %s", prefork)
		app.Logger.Debug().Msgf("Handlers: %d", app.Fiber.HandlersCount())
		app.Logger.Debug().Msgf("Processes: %d", procs)
		app.Logger.Debug().Msgf("PID: %d", os.Getpid())
	}

	// Listen the app (with TLS & HTTP/2 Support)
	var ln net.Listener

	if app.Config.App.TLS.Enable {
		app.Logger.Debug().Msg("TLS support has enabled.")

		cer, err := tls.LoadX509KeyPair(app.Config.App.TLS.CertFile, app.Config.App.TLS.KeyFile)
		if err != nil {
			return err
		}

		cfg := &tls.Config{Certificates: []tls.Certificate{cer}}
		ln, err = tls.Listen("tcp", app.Config.App.Port, cfg)
		if err != nil {
			return err
		}

		if app.Config.App.TLS.HTTP2Support {
			app.Logger.Debug().Msg("HTTP/2 support has enabled.")

			http2.ConfigureServerAndConfig(app.Fiber.Server(), cfg)
		}
	} else {
		ln, err = net.Listen("tcp", app.Config.App.Port)
		if err != nil {
			return err
		}
	}

	err = app.Fiber.Listener(ln)
	if err != nil {
		return err
	}

	return nil
}

// Shutdown the webserver
func (app *AppSkel) Shutdown() {
	// Shutdown fiber
	app.Logger.Warn().Msg("Fiber shutting down.")
	app.Fiber.Shutdown()

	// Shutdown databases
	app.Logger.Warn().Msg("Databases shutting down.")
	app.DB.Redis.Close()
	app.DB.Ent.Close()

	app.Logger.Info().Msgf("%s, was successfully shutted down! \u001b[96mSee you again👋👋\u001b[0m", app.Config.App.Name)
}
