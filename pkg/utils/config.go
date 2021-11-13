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
	}
}

func ParseConfig(config string) (*Config, error) {
	var contents *Config

	_, err := toml.DecodeFile("./config/"+config+".toml", &contents)
	if err != nil {
		return &Config{}, err
	}

	return contents, err
}
