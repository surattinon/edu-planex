package database

import (
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.CurriculumCategories{}, 
		&model.Courses{}, 
		&model.CurriculumCourses{}, 
		&model.Prerequisites{},
		); err != nil {
		return err
	}
	return nil
}
