package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/surattinon/edu-planex/backend/internal/config"
	"github.com/surattinon/edu-planex/backend/internal/database"
	"github.com/surattinon/edu-planex/backend/internal/logger"
	"github.com/surattinon/edu-planex/backend/internal/seeds"
)

func run() {
	cfg, dsn := config.Load()
	logger.Init(true)

	log.Logger.Debug().Msgf("Server Port: %v", cfg.ServerPort)
	log.Logger.Debug().Msgf("DB Host: %v", cfg.DB_Host)
	log.Logger.Debug().Msgf("DB User: %v", cfg.DB_User)
	log.Logger.Debug().Msgf("DB Pass: %v", cfg.DB_Pass)
	log.Logger.Debug().Msgf("DB Name: %v", cfg.DB_Name)
	log.Logger.Debug().Msgf("DB Port: %v", cfg.DB_Port)
	log.Logger.Debug().Msgf("DB SSL Mode: %v", cfg.DB_SSLMode)
	log.Logger.Debug().Msgf("DB DSN: %v", dsn)

	db, err := database.Connect(dsn)
	if err != nil {
		log.Logger.Error().Msgf("%v", err)
		return
	}

	if err := seeds.Load(db, "./internal/seeds/seeds-json/bsc-it-21.json"); err != nil {
		log.Logger.Error().Msgf("%v", err)
	}
	r := gin.Default()
	// r.GET("/courses", handler.GetCourses)
	// r.GET("/course/:code", handler.GetCourseByID)

	port := fmt.Sprintf("localhost:%v", cfg.ServerPort)
	r.Run(port)
}
