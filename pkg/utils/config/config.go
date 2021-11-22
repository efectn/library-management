package config

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Config struct {
	App struct {
		Port       int  `toml:"port"`
		Production bool `toml:"production"`
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
