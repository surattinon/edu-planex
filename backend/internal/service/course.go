package service

import (
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type CourseService struct {
	db *gorm.DB
}

func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{db: db}
}

func (s *CourseService) ListAllCourses() ([]model.Course, error) {
	var cs []model.Course
	if err := s.db.
		Preload("Categories").
		Preload("Prerequisites").
		Find(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}

func (s *CourseService) FindCourseByCode(code string) (*model.Course, error) {
	var course model.Course
	if err := s.db.
		Preload("Categories").
		Preload("Prerequisites").
		First(&course, "course_code = ?", code).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
