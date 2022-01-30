package config

import (
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/rs/zerolog"
)

type app = struct {
	Name        string        `toml:"name"`
	Port        string        `toml:"port"`
	PrintRoutes bool          `toml:"print-routes"`
	Prefork     bool          `toml:"prefork"`
	Production  bool          `toml:"production"`
	IdleTimeout time.Duration `toml:"idle-timeout"`
	TLS         struct {
		Enable       bool
		HTTP2Support bool   `toml:"http2-support"`
		CertFile     string `toml:"cert-file"`
		KeyFile      string `toml:"key-file"`
	}
	Hash struct {
		BcryptCost int `toml:"bcrypt-cost"`
	}
}

type logger = struct {
	TimeFormat string        `toml:"time-format"`
	Level      zerolog.Level `toml:"level"`
	Prettier   bool          `toml:"prettier"`
}

type db = struct {
	Redis struct {
		Url   string `toml:"url"`
		Reset bool   `toml:"reset"`
	}

	Postgres struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Name     string `toml:"name"`
		User     string `toml:"user"`
		Password string `toml:"password"`
	}
}

type middleware = struct {
	Jwt struct {
		Key   string `toml:"secret"`
		Hours time.Duration
	}

	Compress struct {
		Enable bool
		Level  compress.Level
	}

	Recover struct {
		Enable bool
	}

	Monitor struct {
		Enable bool
		Path   string
	}
}

type Config struct {
	App        app
	Logger     logger
	DB         db
	Middleware middleware
}

func ParseConfig(filename string, debug ...bool) (*Config, error) {
	var contents *Config
	var err error

	if len(debug) > 0 {
		_, err = toml.DecodeFile(filename, &contents)
	} else {
		_, err = toml.DecodeFile("./config/"+filename+".toml", &contents)
	}
	if err != nil {
		return &Config{}, err
	}

	return contents, err
}

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
}

// ParseAddr From https://github.com/gofiber/fiber/blob/master/helpers.go#L305.
func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
