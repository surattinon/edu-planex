package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

func InitDB(dsn string) (*gorm.DB, error) {
	log.Info().Msg("Connecting to DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Info().Msg("DB Connected")
	return db, nil
}

func MigrateDB(migrationsDir, url string) error {
	log.Info().Msg("Migrating DB")
	m, err := migrate.New(
		"file://"+migrationsDir,
		url,
	)
	if err != nil {
		return err
	}

	// if err := m.Force(1); err != nil {
	// 	return err
	// }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Info().Msg("DB Migrated")
	return nil
}
