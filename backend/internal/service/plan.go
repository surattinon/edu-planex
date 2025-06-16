package service

import (
	"errors"

	"github.com/surattinon/edu-planex/backend/internal/dto"
	"github.com/surattinon/edu-planex/backend/internal/model"
	"gorm.io/gorm"
)

type PlanService struct {
	db *gorm.DB
}

func NewPlanService(db *gorm.DB) *PlanService {
	return &PlanService{db: db}
}

// List all plans for the single user, preload courses
func (s *PlanService) ListPlans() (*[]model.Plan, error) {
	var plans []model.Plan
	if err := s.db.
		Preload("User").
		Preload("Courses").
		Find(&plans).Error; err != nil {
		return nil, err
	}
	return &plans, nil
}

// Get one plan by ID, with its courses
func (s *PlanService) GetPlan(id uint) (*model.Plan, error) {
	var p model.Plan
	if err := s.db.
		Preload("User").
		Preload("Courses").
		First(&p, "plan_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *PlanService) GetAllPlanTables() (*[]dto.PlanTables, error) {
	var plans []model.Plan
	// preload Courses relationship
	if err := s.db.Preload("Courses").Find(&plans).Error; err != nil {
		return nil, err
	}

	// map to DTO
	out := make([]dto.PlanTables, 0, len(plans))
	for _, p := range plans {
		pr := dto.PlanTables{
			PlanID:   p.PlanID,
			PlanName: p.Name,
			UserID:   p.UserID,
			Courses:  make([]dto.CourseTable, len(p.Courses)),
		}
		for i, crs := range p.Courses {
			pr.Courses[i] = dto.CourseTable{
				CourseCode: crs.CourseCode,
				CourseName: crs.CourseName,
				Credits:    crs.Credits,
			}
		}
		out = append(out, pr)
	}

	return &out, nil
}

func (s *PlanService) GetPlanTableByID(planID int) (*dto.PlanTables, error) {
	var plan model.Plan
	if err := s.db.Preload("Courses").First(&plan, planID).Error; err != nil {
		return nil, err
	}

	resp := dto.PlanTables{
		PlanID:   plan.PlanID,
		PlanName: plan.Name,
		UserID:   plan.UserID,
		Courses:  make([]dto.CourseTable, len(plan.Courses)),
	}
	for i, crs := range plan.Courses {
		resp.Courses[i] = dto.CourseTable{
			CourseCode: crs.CourseCode,
			CourseName: crs.CourseName,
			Credits:    crs.Credits,
		}
	}

	return &resp, nil
}

func (s *PlanService) ApplyPlan(planID uint, year, semNo int) (*model.Enrollment, error) {
	// 0) resolve semester
	semSvc := NewSemesterService(s.db)
	semID, err := semSvc.GetID(year, semNo)
	if err != nil {
		return nil, err
	}

	// 1) load the plan (with its courses)
	plan, err := s.GetPlan(planID)
	if err != nil {
		return nil, err
	}

	// 2) create the enrollment header
	enroll := model.Enrollment{
		UserID:     plan.UserID,
		SemesterID: semID,
	}
	if err := s.db.Create(&enroll).Error; err != nil {
		return nil, err
	}

	// 3) link each course in the plan
	for _, c := range plan.Courses {
		link := model.EnrollmentCourse{
			EnrollmentID: enroll.EnrollmentID,
			CourseCode:   c.CourseCode,
		}
		if err := s.db.Create(&link).Error; err != nil {
			return nil, err
		}
	}

	// 4) reload with preloads for response
	if err := s.db.
		Preload("Semester").
		Preload("Courses").
		First(&enroll, enroll.EnrollmentID).Error; err != nil {
		return nil, err
	}
	return &enroll, nil
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

func (s *PlanService) DeletePlan(id uint) error {
	if err := s.db.
		Where("plan_id = ?", id).
		Delete(&model.Plan{}).Error; err != nil {
		return err
	}
	return nil
}

////////

func (s *PlanService) ListPlansWithApply(userID uint) ([]dto.PlanListDTO, error) {
  // 1) Load all plans + their courses
  var plans []model.Plan
  if err := s.db.
      Where("user_id = ?", userID).
      Preload("Courses", func(db *gorm.DB) *gorm.DB {
        return db.Select("course_code", "course_name", "credits")
      }).
      Find(&plans).Error; err != nil {
    return nil, err
  }

  // 2) Load ALL enrollments for that user
  var enrolls []model.Enrollment
  if err := s.db.
      Where("user_id = ?", userID).
      Preload("Courses", func(db *gorm.DB) *gorm.DB {
        return db.Select("course_code")
      }).
      Find(&enrolls).Error; err != nil {
    return nil, err
  }

  // 3) Build a set of every course_code the user has ever enrolled
  enrolled := make(map[string]struct{})
  for _, e := range enrolls {
    for _, c := range e.Courses {
      enrolled[c.CourseCode] = struct{}{}
    }
  }

  // 4) Build the DTO list
  out := make([]dto.PlanListDTO, len(plans))
  for i, p := range plans {
    var allEnrolled = true
    courses := make([]dto.CourseInfo, len(p.Courses))
    for j, c := range p.Courses {
      courses[j] = dto.CourseInfo{
        CourseCode: c.CourseCode,
        CourseName: c.CourseName,
        Credits:    c.Credits,
      }
      if _, ok := enrolled[c.CourseCode]; !ok {
        allEnrolled = false
      }
    }

    out[i] = dto.PlanListDTO{
      PlanID:  p.PlanID,
      Name:    p.Name,
      UserID:  p.UserID,
      IsApply: allEnrolled,
      Courses: courses,
    }
  }
  return out, nil
}
