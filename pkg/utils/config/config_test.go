package config

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var apiConfig = Config{
	App: app{
		Name:        "LMS app",
		Port:        ":8080",
		Prefork:     false,
		Production:  false,
		IdleTimeout: time.Duration(5),
		TLS: struct {
			Enable       bool
			HTTP2Support bool   `toml:"http2-support"`
			CertFile     string `toml:"cert-file"`
			KeyFile      string `toml:"key-file"`
		}{
			Enable:       false,
			HTTP2Support: false,
			CertFile:     "./storage/selfsigned.crt",
			KeyFile:      "./storage/selfsigned.key",
		},
		Hash: struct {
			BcryptCost int `toml:"bcrypt-cost"`
		}{BcryptCost: 10},
	},
	Logger: logger{
		TimeFormat: "",
		Level:      zerolog.Level(0),
		Prettier:   true,
	},
	DB: db{
		Postgres: struct {
			Host     string `toml:"host"`
			Port     int    `toml:"port"`
			Name     string `toml:"name"`
			User     string `toml:"user"`
			Password string `toml:"password"`
		}{
			Host:     "postgres",
			Port:     5432,
			Name:     "library_management",
			User:     "postgres",
			Password: "postgres",
		},
		Redis: struct {
			Url   string `toml:"url"`
			Reset bool   `toml:"reset"`
		}{
			Url:   "redis://redis:6379/",
			Reset: false,
		},
	},
	Middleware: middleware{
		Jwt: struct {
			Key   string `toml:"secret"`
			Hours time.Duration
		}{
			Key:   "UjXn2r5u8x/A?D(G+KbPeSgVkYp3s6v9",
			Hours: 5,
		},
		Compress: struct {
			Enable bool
			Level  compress.Level
		}{
			Enable: true,
			Level:  compress.Level(1),
		},
		Recover: struct{ Enable bool }{
			Enable: true,
		},
		Monitor: struct {
			Enable bool
			Path   string
		}{
			Enable: true,
			Path:   "/monitor",
		},
		Pprof: struct{ Enable bool }{
			Enable: false,
		},
	},
}

func Test_ParseConfig(t *testing.T) {
	t.Parallel()

	config, err := ParseConfig("../../../config/api.toml", true)

	assert.Equal(t, nil, err)
	assert.Equal(t, apiConfig, *config)
}

func Benchmark_ParseConfig(b *testing.B) {
	var err error

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = ParseConfig("../../../config/api.toml", true)
	}

	assert.Equal(b, nil, err)
}

func Test_ParseAddr(t *testing.T) {
	testCases := []struct {
		addr, host, port string
	}{
		{"[::1]:3000", "[::1]", "3000"},
		{"127.0.0.1:3000", "127.0.0.1", "3000"},
		{"/path/to/unix/socket", "/path/to/unix/socket", ""},
	}

	for _, c := range testCases {
		host, port := ParseAddr(c.addr)
		assert.Equal(t, c.host, host, "addr host")
		assert.Equal(t, c.port, port, "addr port")
	}
}

/*
func Test_IsEnabled(t *testing.T) {
	ifTrue := IsEnabled(true)
	ifFalse := IsEnabled(false)

	assert.Equal(t, func(c *fiber.Ctx) bool { return false }, ifTrue)
	assert.Equal(t, func(c *fiber.Ctx) bool { return true }, ifFalse)
}
*/
