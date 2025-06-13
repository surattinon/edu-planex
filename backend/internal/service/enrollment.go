package service

import (
	"github.com/surattinon/edu-planex/backend/internal/dto"
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

func (s *EnrollService) GetHistory(userID uint) ([]dto.EnrollmentHistoryDTO, error) {
    // 1) Load enrollments with Semester and Courses (code, name, credits)
    var enrolls []model.Enrollment
    if err := s.db.
        Where("user_id = ?", userID).
        Preload("Semester").
        Preload("Courses", func(db *gorm.DB) *gorm.DB {
            return db.Select("course_code", "course_name", "credits")
        }).
        Order("semester_id DESC").
        Find(&enrolls).Error; err != nil {
        return nil, err
    }

    // 2) Map to DTOs
    result := make([]dto.EnrollmentHistoryDTO, len(enrolls))
    for i, e := range enrolls {
        dtoCourses := make([]dto.CourseInfo, len(e.Courses))
        for j, c := range e.Courses {
            dtoCourses[j] = dto.CourseInfo{
                CourseCode: c.CourseCode,
                CourseName: c.CourseName,
                Credits:    c.Credits,
            }
        }
        result[i] = dto.EnrollmentHistoryDTO{
            EnrollmentID: e.EnrollmentID,
            UserID:       e.UserID,
            Semester: dto.SemesterInfo{
                Year:   e.Semester.Year,
                Number: e.Semester.SemesterNo,
            },
            Courses: dtoCourses,
        }
    }
    return result, nil
}
