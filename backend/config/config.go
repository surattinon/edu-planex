package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server_Port   string
	DB_Url        string
	DB_Dsn        string
	DB_Host       string
	DB_User       string
	DB_Pass       string
	DB_Name       string
	DB_Port       string
	DB_SSLMode    string
	Migration_Dir string
	JWT_Secret    string
	Debug         bool
}

func Load() *Config {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // restructure ENV string

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
		return nil
	}

	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%v/%s?sslmode=%s",
		v.GetString("db.user"),
		v.GetString("db.pass"),
		v.GetString("db.host"),
		v.GetString("db.port"),
		v.GetString("db.name"),
		v.GetString("db.sslmode"),
	)

	dbDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		v.GetString("db.host"),
		v.GetString("db.port"),
		v.GetString("db.user"),
		v.GetString("db.pass"),
		v.GetString("db.name"),
		v.GetString("db.sslmode"),
	)

	cfg := Config{
		Server_Port:   v.GetString("server.port"),
		DB_Url:        dbURL,
		DB_Dsn:        dbDSN,
		DB_Host:       v.GetString("db.host"),
		DB_User:       v.GetString("db.user"),
		DB_Pass:       v.GetString("db.pass"),
		DB_Name:       v.GetString("db.name"),
		DB_Port:       v.GetString("db.port"),
		DB_SSLMode:    v.GetString("db.sslmode"),
		Migration_Dir: v.GetString("migrations.dir"),
		JWT_Secret:    v.GetString("jwt.secret"),
		Debug:         v.GetBool("server.debug"),
	}

	log.Print("Config loaded")
	return &cfg
}
