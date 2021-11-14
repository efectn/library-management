package utils

import "github.com/BurntSushi/toml"

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
}

func ParseConfig(filename string) (*Config, error) {
	var contents *Config

	_, err := toml.DecodeFile("./config/"+filename+".toml", &contents)
	if err != nil {
		return &Config{}, err
	}

	return contents, err
}
