package service

import (
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
		PlanID:  plan.PlanID,
		PlanName: plan.Name,
		UserID:  plan.UserID,
		Courses: make([]dto.CourseTable, len(plan.Courses)),
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
