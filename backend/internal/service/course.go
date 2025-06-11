package service

import (
	"github.com/surattinon/edu-planex/backend/internal/dto"
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

func (s *CourseService) GetCourseTable() ([]dto.CourseTable, error) {
	var table []dto.CourseTable
	if err := s.db.
		Model(&model.Course{}).
		Select("course_code", "course_name", "credits").
		Find(&table).Error; err != nil {
		return nil, err
	}
	return table, nil
}

func (s *CourseService) ListSummariesWithPrereqs() ([]dto.CourseTable, error) {
	// 1) Load minimal course fields + loaded prerequisites
	var courses []model.Course
	if err := s.db.
		Select("course_code", "course_name", "credits").
		Preload("Prerequisites").
		Find(&courses).Error; err != nil {
		return nil, err
	}

	// 2) Map into summaries
	out := make([]dto.CourseTable, len(courses))
	for i, c := range courses {
		codes := make([]string, len(c.Prerequisites))
		for j, p := range c.Prerequisites {
			codes[j] = p.PreCourseCode // your struct field
		}
		out[i] = dto.CourseTable{
			CourseCode:    c.CourseCode,
			CourseName:    c.CourseName,
			Credits:       c.Credits,
			Prerequisites: codes,
		}
	}
	return out, nil
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
