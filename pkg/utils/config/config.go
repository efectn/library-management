package config

import (
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/rs/zerolog"
)

type Config struct {
	App struct {
		Name       string `toml:"name"`
		Port       string `toml:"port"`
		Prefork    bool   `toml:"prefork"`
		Production bool   `toml:"production"`
	}

	Logger struct {
		TimeFormat string        `toml:"time-format"`
		Level      zerolog.Level `toml:"level"`
		Prettier   bool          `toml:"prettier"`
	}

	DB struct {
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

	Middleware struct {
		Jwt struct {
			Key   string
			Hours time.Duration
		}

		Compress struct {
			Enable bool
			Level  compress.Level
		}

		Recover struct {
			Enable bool
		}
	}
}

func ParseConfig(filename string) (*Config, error) {
	var contents *Config

	_, err := toml.DecodeFile("./config/"+filename+".toml", &contents)
	if err != nil {
		return &Config{}, err
	}

	return contents, err
}

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	enabled := true
	if key {
		enabled = false
	}

	return func(c *fiber.Ctx) bool { return enabled }
}

// From https://github.com/gofiber/fiber/blob/master/helpers.go#L305.
func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
