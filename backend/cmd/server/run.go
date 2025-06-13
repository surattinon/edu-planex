package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/surattinon/edu-planex/backend/internal/config"
	"github.com/surattinon/edu-planex/backend/internal/database"
	"github.com/surattinon/edu-planex/backend/internal/handler"
	"github.com/surattinon/edu-planex/backend/internal/logger"
	"github.com/surattinon/edu-planex/backend/internal/service"
)

func run() {
	// Load Config
	cfg, dsn := config.Load()

	// Init Logger
	logger.Init(false)
	log.Logger.Debug().Msgf("Server Port: %v", cfg.ServerPort)
	log.Logger.Debug().Msgf("DB Host: %v", cfg.DB_Host)
	log.Logger.Debug().Msgf("DB User: %v", cfg.DB_User)
	log.Logger.Debug().Msgf("DB Pass: %v", cfg.DB_Pass)
	log.Logger.Debug().Msgf("DB Name: %v", cfg.DB_Name)
	log.Logger.Debug().Msgf("DB Port: %v", cfg.DB_Port)
	log.Logger.Debug().Msgf("DB SSL Mode: %v", cfg.DB_SSLMode)
	log.Logger.Debug().Msgf("DB DSN: %v", dsn)

	// Connect to Postgres
	db, err := database.Connect(dsn)
	if err != nil {
		log.Logger.Error().Msgf("%v", err)
		return
	}

	// New Services
	crsSvc := service.NewCourseService(db)
	catSvc := service.NewCategoryService(db)
	userSvc := service.NewUserService(db)
	planSvc := service.NewPlanService(db)
	enrollSvc := service.NewEnrollService(db)
	progSvc := service.NewProgressService(db)
	curSvc := service.NewCurriculumService(db)

	// New Handlers
	crsHnd := handler.NewCourseHandler(crsSvc)
	catHnd := handler.NewCategoryHandler(catSvc)
	userHnd := handler.NewUserHandler(userSvc)
	planHnd := handler.NewPlanHandler(planSvc)
	enrollHnd := handler.NewEnrollHandler(enrollSvc)
	progHnd := handler.NewProgressHandler(progSvc)
	curHnd := handler.NewCurriculumHandler(curSvc)

	// init API router
	r := gin.Default()

	// GET
	r.GET("/courses", crsHnd.GetCourseList)
	r.GET("/coursetable", crsHnd.CourseTable)
	r.GET("/course/:code", crsHnd.GetCourseByCode)
	r.GET("/course/:code/categories", crsHnd.GetCatByCode)
	r.GET("/categories", catHnd.GetCatList)
	r.GET("/categories/:id", catHnd.GetCatByID)
	r.GET("/profile", userHnd.GetProfile)
	r.GET("/plans", planHnd.GetPlanList)
	r.GET("/plan/:id", planHnd.GetPlanByID)
	r.GET("/plantable/:id", planHnd.PlanTable)
	r.GET("/plantable", planHnd.AllPlanTable)
	r.GET("/enrollments", enrollHnd.GetEnrollList)
	r.GET("/enrollhistory", enrollHnd.GetEnrollBySemester)
	r.GET("/enrollyear", enrollHnd.GetEnrollByYear)
	r.GET("/progress", progHnd.GetProgress)

	// POST
	r.POST("/plan/:id/apply", planHnd.Apply)
	r.POST("/plans", planHnd.Create)

	// DELETE
	r.DELETE("/plan/:id", planHnd.Delete)

	// PUT
	r.PUT("/profile", userHnd.UpdateProfile)

	v1 := r.Group("api/v1")
	{
		v1.GET("/progress", progHnd.GetProgress)
		v1.GET("/enrollhist", enrollHnd.EnrollHistoryList)
		v1.GET("/plantable", planHnd.AllPlanTable)
		v1.GET("/curriculum", crsHnd.GetCourseList)
		v1.GET("/curriculumtable", curHnd.List)

		v1.POST("/plan/:id/apply", planHnd.Apply)
		v1.POST("/plans", planHnd.Create)

		v1.PUT("/profile", userHnd.UpdateProfile)

		v1.DELETE("/plan/:id", planHnd.Delete)
	}

	// Serve
	r.Run(":" + cfg.ServerPort)
}
