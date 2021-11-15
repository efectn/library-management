package utils

import (
	"time"

	"github.com/BurntSushi/toml"
)

type ConfigBase struct {
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
	}
}

var Config = new(ConfigBase)

func ParseConfig(filename string) (*ConfigBase, error) {
	var contents *ConfigBase

	_, err := toml.DecodeFile("./config/"+filename+".toml", &contents)
	if err != nil {
		return &ConfigBase{}, err
	}

	return contents, err
}
