package database

import (
	"context"
	"fmt"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/storage/s3"
	"github.com/rs/zerolog"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Ent   *ent.Client
	Redis *redis.Storage
	S3    *s3.Storage
}

type Seeder interface {
	Seed() error
	Count() (int, error)
}

func Init() *Database {
	return new(Database)
}

func (db *Database) SetupEnt(host string, port int, user string, password string, name string, logger ...zerolog.Logger) error {
	conn, err := sql.Open("pgx", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, name))
	if err != nil {
		return err
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)

	// Setup Logger
	/*if len(logger) > 0 {
		drvv := dialect.DebugWithContext(drv, func(ctx context.Context, args ...interface{}) {
			op := drv.
				logger[0].Debug().Msgf("entgo: query=%v args=%v", op.Query, args)
		})
	}*/
	/*drvv := dialect.DebugWithContext(drv, func(ctx context.Context, args ...interface{}) {
		logger[0].Debug().Msgf("entgo: query=%v args=%v", "op.Query", args)
	})*/
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

func (db *Database) SetupS3(endpoint, bucket, region, accessKey, secretKey string) error {
	conn := s3.New(s3.Config{
		Bucket:   bucket,
		Region:   region,
		Endpoint: endpoint,
		Credentials: s3.Credentials{
			AccessKey:       accessKey,
			SecretAccessKey: secretKey,
		},
	})

	db.S3 = conn

	return nil
}

func (db *Database) MigrateModels() error {
	if err := db.Ent.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err.Error())
	}

	return nil

}

func (db *Database) SeedModels(logger zerolog.Logger, seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			logger.Panic().Err(err).Msg("")
		}

		if count == 0 {
			err = v.Seed()
			if err != nil {
				logger.Panic().Err(err).Msg("")
			}

			logger.Debug().Msg("Table has seeded successfully.")
		} else {
			logger.Warn().Msg("Table has seeded already. Skipping!")
		}
	}
}
