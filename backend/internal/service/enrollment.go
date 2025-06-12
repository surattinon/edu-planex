package service

import (
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type EnrollService struct {
	db *gorm.DB
}

func NewEnrollService(db *gorm.DB) *EnrollService {
	return &EnrollService{db: db}
}

func (s *EnrollService) ListEnroll() ([]model.Enrollment, error) {
	var list []model.Enrollment
	if err := s.db.
		Preload("Semester").
		Preload("Courses").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *EnrollService) ListBySemester(year, semNo int) ([]model.Enrollment, error) {
	var list []model.Enrollment
	err := s.db.
		Model(&model.Enrollment{}). // specify base table
		Joins("JOIN semesters ON enrollments.semester_id = semesters.semester_id").
		Where("semesters.year = ? AND semesters.semester_number = ?", year, semNo).
		Preload("Semester").
		Preload("Courses").
		Find(&list).Error
	return list, err
}

func (s *EnrollService) ListByYear(year int) ([]model.Enrollment, error) {
	var list []model.Enrollment
	err := s.db.
		Model(&model.Enrollment{}). // specify base table
		Joins("JOIN semesters ON enrollments.semester_id = semesters.semester_id").
		Where("semesters.year = ? ", year).
		Preload("Semester").
		Preload("Courses").
		Find(&list).Error
	return list, err
}

// func (s *EnrollService) GetCredsByID(id int) (*dto.CreditResult, error) {
// 	var creds dto.CreditResult
// }
