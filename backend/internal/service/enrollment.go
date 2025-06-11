package service

import (
	"errors"

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

func (s *PlanService) CreatePlan(userID uint, name string, codes []string) (*model.Plan, error) {
	if len(codes) == 0 {
		return nil, errors.New("must specify at least one course")
	}

	// wrap in a transaction for atomicity
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Plan header, link to user info
	plan := model.Plan{
		UserID: userID,
		Name:   name,
	}
	if err := tx.Create(&plan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create each plan_course link
	links := make([]model.PlanCourse, 0, len(codes))
	for _, code := range codes {
		links = append(links, model.PlanCourse{
			PlanID:     plan.PlanID,
			CourseCode: code,
		})
	}
	if err := tx.Create(&links).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Preload the courses and return created plan
	if err := tx.
		Preload("Courses").
		First(&plan, plan.PlanID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &plan, nil
}
