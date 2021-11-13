package types

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
