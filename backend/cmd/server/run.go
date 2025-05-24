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

	db, err := database.InitDB(dbDSN)
	if err != nil {
		log.Error().Msgf("Could not connect to DB: %v", err)
	}
	if err := database.MigrateDB(cfg.Migration_Dir, dbUrl); err != nil {
		log.Error().Msgf("Could not migrate DB: %v", err)
	}

	// Repos & services
	userRepo := repository.NewUserRepo(db)
	authService := service.NewAuthService(userRepo, cfg.JWT_Secret)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)

	// Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api/v1")
	auth := api.Group("/auth") 
	{
		auth.POST("/advisor/signup", authHandler.AdvisorSignup)
		auth.POST("/login", authHandler.Login)
	}

	protected := api.Group("/")
	protected.Use(middleware.JWTMiddleware(cfg.JWT_Secret))
	{
		protected.POST("/student/signup", authHandler.StudentSignup)
	}

	port := cfg.Server_Port
	log.Info().Msgf("Start server on port: %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
