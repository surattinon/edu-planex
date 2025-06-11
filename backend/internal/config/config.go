package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerPort string    `env:"SERVER_PORT" envDefault:"8888"`
	DB_Host    string `env:"DB_HOST" envDefault:"localhost"`
	DB_User    string `env:"DB_USER" envDefault:"postgres"`
	DB_Pass    string `env:"DB_PASS" envDefault:"eduplanex442"`
	DB_Name    string `env:"DB_NAME" envDefault:"eduplanex"`
	DB_Port    int    `env:"DB_PORT" envDefault:"5432"`
	DB_SSLMode string `env:"DB_SSL_MODE" envDefault:"disable"`
	TimeZone string `env:"TIME_ZONE" envDefault:"Asia/Bangkok"`
}

func Load() (Config, string) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	dburl := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v",
		cfg.DB_Host,
		cfg.DB_Port,
		cfg.DB_User,
		cfg.DB_Pass,
		cfg.DB_Name,
		cfg.DB_SSLMode,
		cfg.TimeZone,
	)

	return cfg, dburl
}
