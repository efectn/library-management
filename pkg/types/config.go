package types

type Config struct {
	App struct {
		Port       int  `toml:"port"`
		Production bool `toml:"production"`
	}
}
