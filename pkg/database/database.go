package database

import (
	"github.com/gofiber/storage/redis"
	"gorm.io/gorm"
)

type Database struct {
	Gorm  *gorm.DB
	Redis *redis.Storage
}

func Init() *Database {
	return new(Database)
}

/*func (db *Database) SetupGORM() error {
	conn, err := gorm.Open(postgres.Open("---"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Gorm = conn

	return nil
}*/

func (db *Database) SetupRedis(url string, reset bool) error {
	conn := redis.New(redis.Config{
		URL:   url,
		Reset: reset,
	})

	db.Redis = conn

	return nil
}
