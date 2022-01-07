package database

import (
	"context"
	"fmt"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/gofiber/storage/redis"
	"github.com/rs/zerolog/log"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Ent   *ent.Client
	Redis *redis.Storage
}

type Seeder interface {
	Seed() error
	Count() (int, error)
}

func Init() *Database {
	return new(Database)
}

func (db *Database) SetupEnt(host string, port int, user string, password string, name string) error {
	conn, err := sql.Open("pgx", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, name))
	if err != nil {
		return err
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)
	db.Ent = ent.NewClient(ent.Driver(drv))

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

func (db *Database) MigrateModels() error {
	if err := db.Ent.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err.Error())
	}

	return nil

}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			log.Panic().Err(err).Msg("")
		}

		if count == 0 {
			err = v.Seed()
			if err != nil {
				log.Panic().Err(err).Msg("")
			}
		} else {
			log.Warn().Msg("Table has seeded already. Skipping!")
		}
	}
}
