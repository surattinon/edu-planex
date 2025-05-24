package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/surattinon/edu-planex/backend/config"
	"github.com/surattinon/edu-planex/backend/internal/database"
	"github.com/surattinon/edu-planex/backend/internal/handler"
	"github.com/surattinon/edu-planex/backend/internal/logger"
	"github.com/surattinon/edu-planex/backend/internal/middleware"
	"github.com/surattinon/edu-planex/backend/internal/repository"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

func run() {
	// Load Config
	cfg := config.Load()

	// Init Logger
	debugMode := cfg.Debug
	logger.Init(debugMode)

	log.Info().Msg("Server starting...")

	// DB connection and migration
	dbUrl := cfg.DB_Url
	dbDSN := cfg.DB_Dsn

	log.Printf("DB URL is: %s", dbUrl)
	log.Printf("DB DSN is: %s", dbDSN)
}
