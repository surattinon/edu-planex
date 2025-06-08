package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/surattinon/edu-planex/backend/internal/config"
	"github.com/surattinon/edu-planex/backend/internal/database"
	"github.com/surattinon/edu-planex/backend/internal/handler"
	"github.com/surattinon/edu-planex/backend/internal/logger"
	"github.com/surattinon/edu-planex/backend/internal/service"
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

	crsSvc := service.NewCourseService(db)
	crsHnd := handler.NewCourseHandler(crsSvc)
	catSvc := service.NewCategoryService(db)
	catHnd := handler.NewCategoryHandler(catSvc)

	r := gin.Default()
	r.GET("/courses", crsHnd.GetCourseList)
	r.GET("/course/:code", crsHnd.GetCourseByCode)
	r.GET("/course/:code/categories", crsHnd.GetCatByCode)
	r.GET("/categories", catHnd.GetCatList)
	r.GET("/categories/:id", catHnd.GetCatByID)

	port := fmt.Sprintf("localhost:%v", cfg.ServerPort)
	r.Run(port)
}
