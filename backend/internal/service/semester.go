package service

import (
	"errors"

	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type SemesterService struct {
	db *gorm.DB
}

func NewSemesterService(db *gorm.DB) *SemesterService {
	return &SemesterService{db: db}
}

// GetID returns the semester_id for the given year & number, or error if not found.
func (s *SemesterService) GetID(year, semNo int) (uint, error) {
	var sem model.Semester
	err := s.db.
		Where("year = ? AND semester_number = ?", year, semNo).
		First(&sem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("semester not found; please seed it in semesters table")
	}
	return sem.SemesterID, err
}
