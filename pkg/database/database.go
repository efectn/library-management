package database

import (
	"fmt"

	"github.com/gofiber/storage/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm  *gorm.DB
	Redis *redis.Storage
}

type Seeder interface {
	Seed()
	ReturnModel() interface{}
}

func Init() *Database {
	return new(Database)
}

func (db *Database) SetupGORM(host string, port int, user string, password string, name string) error {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, user, password, name)), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Gorm = conn

	return nil
}

func (db *Database) SetupRedis(url string, reset bool) error {
	conn := redis.New(redis.Config{
		URL:   url,
		Reset: reset,
	})

	db.Redis = conn

	return nil
}

func (db *Database) MigrateModels(models ...interface{}) error {
	return db.Gorm.AutoMigrate(models...)
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {
		var count int64 = 0
		db.Gorm.Model(v.ReturnModel()).Count(&count)

		if count == 0 {
			v.Seed()
		} else {
			fmt.Println("=====> WARN: Table has seeded already. Skipping!")
		}
	}
}
